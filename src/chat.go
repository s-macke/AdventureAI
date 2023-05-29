package zmachine

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"regexp"
	"strings"
)

const (
	InfoColor = "\033[1;34m%s\033[0m"
)

const systemMsg string = `You act as a player of an interactive text adventure. The goal is to win the game. 
The user provides the text of the text adventure. He is not a human and just prints the output of the game.

The format of your output must be:
NARRATIVE: A short description of the current narrative you are in
THOUGHT: Your thought about the situation and what to do next
COMMAND: The command you want to execute. Must always begin with a verb. The maximum number of words are 4. The commands should be very simple.

Your name is not Brian Hadley. You have accidentally killed Brian Hadley in the house.
Your first task is to look under your bed.`

// You are a murderer.

type Command struct {
	Situation string `json:"situation"`
	Thought   string `json:"thought"`
	Command   string `json:"command"`
}

type StoryStep struct {
	IsResponse       bool                           `json:"isResponse"`
	Command          Command                        `json:"command"`
	CompletionTokens int                            `json:"CompletionTokens"`
	PromptTokens     int                            `json:"PromptTokens"`
	Messages         []openai.ChatCompletionMessage `json:"messages"`
}

type StoryState struct {
	Prompt string      `json:"prompt"`
	Steps  []StoryStep `json:"steps"`
}

func StoreStoryToFile(state *StoryState, name string) {
	stateAsJson, err := json.MarshalIndent(state, "", " ")
	if err != nil {
		panic(err)
	}
	filename := "storydump/" + name + ".json"
	err = os.WriteFile(filename, stateAsJson, 0644)
	if err != nil {
		panic(err)
	}
}

func LoadStoryFromFile(state *StoryState, name string) {
	filename := "storydump/" + name + ".json.backup"
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &state)
	if err != nil {
		panic(err)
	}
}

type ChatState struct {
	client                *openai.Client
	messages              []openai.ChatCompletionMessage
	zm                    *ZMachine
	output                string
	totalCompletionTokens int
	totalPromptTokens     int
	story                 StoryState
	currentStoryStep      int
}

func NewChatState(zm *ZMachine) *ChatState {
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		panic(" OPENAI_API_KEY env var not set")
	}
	cs := &ChatState{
		client: openai.NewClient(key),
		story: StoryState{
			Prompt: systemMsg,
		},
		totalPromptTokens:     0,
		totalCompletionTokens: 0,
		output:                "",
		zm:                    zm,
		currentStoryStep:      0,
	}
	cs.zm.input = cs.chatInput
	cs.messages = append(cs.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemMsg,
	})

	//LoadStoryFromFile(&cs.story, cs.zm.name)

	return cs
}

func separateCommand(content string) Command {
	cmd := Command{}

	if content == "" {
		panic("empty content")
	}

	re := regexp.MustCompile(`\r?\n`)
	content = re.ReplaceAllString(content, " ")
	re = regexp.MustCompile(`(.*)THOUGHT:(.*)COMMAND:(.*)`)
	matches := re.FindStringSubmatch(content)

	cmd.Situation = strings.TrimSpace(matches[1])
	cmd.Thought = strings.TrimSpace(matches[2])
	cmd.Command = strings.TrimSpace(matches[3])
	cmd.Command = strings.ReplaceAll(cmd.Command, ".", "")

	if cmd.Command == "" {
		panic("empty command")
	}

	return cmd
}

func (cs *ChatState) chatInput() string {
	fmt.Print(cs.output)

	cs.messages = append(cs.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: cs.output,
	})
	cs.zm.output.Reset()
	cs.output = ""


	// just return to the previous state if we have commands left
	if cs.currentStoryStep < len(cs.story.Steps) {
		step := cs.story.Steps[cs.currentStoryStep]
		if step.IsResponse {
			cs.messages = step.Messages
			fmt.Printf(InfoColor, "SITUATION: "+step.Messages[len(cs.messages)-1].Content)
			fmt.Println()
			cs.totalCompletionTokens += step.CompletionTokens
			cs.totalPromptTokens += step.PromptTokens
			cs.currentStoryStep++
			return step.Command.Command
		} else {
			panic("not a response")
		}
	}

	cs.messages = append(cs.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: "NARRATIVE: ",
	})

	resp, err := cs.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            openai.GPT4,
			Messages:         cs.messages,
			MaxTokens:        256,
			PresencePenalty:  0,
			FrequencyPenalty: 0,
		},
	)
	cs.totalCompletionTokens += resp.Usage.CompletionTokens
	cs.totalPromptTokens += resp.Usage.PromptTokens
	fmt.Printf("PromptTokens: %d CompletionTokens: %d\n", resp.Usage.PromptTokens, resp.Usage.CompletionTokens)
	fmt.Printf("totalPromptTokens: %d totalCompletionTokens: %d price: %.2f\n", cs.totalPromptTokens, cs.totalCompletionTokens, float32(cs.totalCompletionTokens)*0.00006+float32(cs.totalPromptTokens)*0.00003)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		panic("ChatCompletion error")
	}

	content := resp.Choices[0].Message.Content

	cs.messages = append(cs.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})

	fmt.Printf(InfoColor, "SITUATION: "+content)
	fmt.Println()

	cmd := separateCommand(content)

	cs.story.Steps = append(cs.story.Steps, StoryStep{
		IsResponse:       true,
		Command:          cmd,
		Messages:         cs.messages,
		CompletionTokens: resp.Usage.CompletionTokens,
		PromptTokens:     resp.Usage.PromptTokens,
	})
	StoreStoryToFile(&cs.story, cs.zm.name)
	cs.currentStoryStep++

	fmt.Println("Press ENTER to continue...")
	_, _ = fmt.Scanln()

	return cmd.Command
}

func (cs *ChatState) chatLoop() {
	for !cs.zm.done {
		cs.zm.InterpretInstruction()
		if cs.zm.output.Len() > 0 {
			if cs.zm.windowId == 0 {
				cs.output += cs.zm.output.String()
			}
			cs.zm.output.Reset()
		}
	}
}

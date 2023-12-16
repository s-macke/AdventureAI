package chat

import (
	"fmt"
	"github.com/s-macke/AdventureAI/src/chat/backend"
	"github.com/s-macke/AdventureAI/src/zmachine"
	"regexp"
	"strings"
)

const (
	InfoColor = "\033[1;34m%s\033[0m"
)

const systemMsg string = `You act as a player of an interactive text adventure. The goal is to win the game. 
The user provides the text of the text adventure. He is not a human and just prints the output of the game.

The format of your output must be:
SITUATION: {A short description of the current situation you are in.}
THOUGHT: {A curious, adventurous thought.}
COMMAND: {The single two word command you want to execute.}
`

//Your name is not Brian Hadley. You have accidentally killed Brian Hadley in the house.
//Your first task is to look under your bed.`
// You are a murderer.

type ChatState struct {
	chatClient       backend.ChatBackend
	zm               *zmachine.ZMachine
	story            StoryState
	output           string
	currentStoryStep int
}

func NewChatState(zm *zmachine.ZMachine, backendAsString string) *ChatState {
	cs := &ChatState{
		story: StoryState{
			Prompt: systemMsg,
		},
		output:           "",
		zm:               zm,
		currentStoryStep: 0,
	}
	fmt.Println("Use backend: ", backendAsString)
	switch backendAsString {
	case "gpt3", "gpt4":
		cs.chatClient = backend.NewOpenAIChat(systemMsg, backendAsString)
	case "orca2":
		cs.chatClient = backend.NewLlamaChat(systemMsg, backendAsString)
	case "mistral":
		cs.chatClient = backend.NewMistralChat(systemMsg)
	case "gemini":
		cs.chatClient = backend.NewVertexAIChat(systemMsg)
	default:
		panic("Unknown backend")
	}

	cs.zm.Input = cs.chatInput
	//LoadStoryFromFile(&cs.story, cs.zm.Name)
	return cs
}

func separateCommand(content string) Command {
	cmd := Command{}

	if content == "" {
		panic("empty content")
	}

	re := regexp.MustCompile(`\r?\n`)
	content = re.ReplaceAllString(content, " ")
	re = regexp.MustCompile(`SITUATION:(.*)THOUGHT:(.*)COMMAND:(.*)`)
	matches := re.FindStringSubmatch(content)

	cmd.Situation = strings.TrimSpace(matches[1])
	//cmd.Narrator = strings.TrimSpace(matches[2])
	cmd.Thought = strings.TrimSpace(matches[2])
	cmd.Command = strings.TrimSpace(matches[3])
	if cmd.Command[0] == '"' && cmd.Command[len(cmd.Command)-1] == '"' {
		cmd.Command = cmd.Command[1 : len(cmd.Command)-1]
	}
	cmd.Command = strings.ReplaceAll(cmd.Command, ".", "")
	if cmd.Command == "" {
		panic("empty command")
	}

	return cmd
}

func (cs *ChatState) chatInput() string {
	cs.story.Messages = append(cs.story.Messages, StoryMessage{
		Role:             "user",
		Content:          cs.output,
		CompletionTokens: 0,
		PromptTokens:     0,
		IsResponse:       false,
		Command:          Command{},
	})

	fmt.Print(cs.output)

	// just return to the previous state if we have commands left
	if (cs.currentStoryStep*2 + 1) < len(cs.story.Messages) {
		step := cs.story.Messages[cs.currentStoryStep*2+1]
		if step.IsResponse {
			//fmt.Printf(InfoColor, "REASONING: "+step.Messages[len(cs.messages)-1].Content)
			fmt.Printf(InfoColor, "COMMAND: "+step.Command.Command+"\n")
			cs.currentStoryStep++
			return step.Command.Command
		} else {
			//panic("not a response")
		}
	}

	content, promptTokens, completionTokens := cs.chatClient.GetResponse(cs.output)
	cs.zm.Output.Reset()
	cs.output = ""

	fmt.Printf(InfoColor, content)
	fmt.Println()

	cmd := separateCommand(content)

	cs.story.Messages = append(cs.story.Messages, StoryMessage{
		Role:             "assistant",
		Content:          cs.output,
		CompletionTokens: completionTokens,
		PromptTokens:     promptTokens,
		IsResponse:       true,
		Command:          cmd,
	})

	StoreStoryToFile(&cs.story, cs.zm.Name)
	cs.currentStoryStep++

	if cs.currentStoryStep%5 == 0 {
		fmt.Println("Press ENTER to continue...")
		_, _ = fmt.Scanln()
	}

	return cmd.Command
}

func (cs *ChatState) ChatLoop() {
	for !cs.zm.Done {
		cs.zm.InterpretInstruction()
		if cs.zm.Output.Len() > 0 {
			if cs.zm.WindowId == 0 {
				cs.output += cs.zm.Output.String()
			}
			cs.zm.Output.Reset()
		}
	}
}

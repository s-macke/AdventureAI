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

var client *openai.Client
var messages []openai.ChatCompletionMessage
var zm *ZMachine
var output = ""
var totalCompletionTokens = 0
var totalPromptTokens = 0
var storyStep = 0

func StoreMessages(messages []openai.ChatCompletionMessage, storyStep int, isRequest bool) {
	messagesAsJson, _ := json.MarshalIndent(messages, "", " ")
	filename := ""
	if isRequest {
		filename = fmt.Sprintf("storydump/request_%03d.json", storyStep)
	} else {
		filename = fmt.Sprintf("storydump/response_%03d.json", storyStep)
	}
	err := os.WriteFile(filename, messagesAsJson, 0644)
	if err != nil {
		panic(err)
	}
}

const systemMsg string = `
You act as a player of an interactive text adventure. The goal is to win the game. 
The user provides the text of the text adventure. He is not a human and just prints the output of the game.

The format of your output must be:
SITUATION: A short description of the current situation you are in
THOUGHT: Your thought about the situation and what to do next
COMMAND: The command you want to execute

The commands always start with a verb. Do never enter an empty command line.
.`

//Do not respond with anything at all other than the command that the player would enter.

func chatInput() string {
	_, _ = os.Stdout.WriteString(output)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: output,
	})
	zm.output.Reset()
	output = ""

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: "SITUATION: ",
	})

	StoreMessages(messages, storyStep, true)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            openai.GPT4,
			Messages:         messages,
			MaxTokens:        256,
			PresencePenalty:  0,
			FrequencyPenalty: 0,
		},
	)
	totalCompletionTokens += resp.Usage.CompletionTokens
	totalPromptTokens += resp.Usage.PromptTokens
	fmt.Printf("%d %d %.2f\n", totalPromptTokens, totalCompletionTokens, float32(totalCompletionTokens)*0.00006+float32(totalPromptTokens)*0.00003)
	//fmt.Printf("ChatCompletion response: %v\n", resp)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		panic("ChatCompletion error")
	}

	content := resp.Choices[0].Message.Content

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
	StoreMessages(messages, storyStep, false)

	fmt.Printf(InfoColor, content)
	fmt.Println()
	_, _ = fmt.Scanln()

	re := regexp.MustCompile(`\r?\n`)
	content = re.ReplaceAllString(content, " ")
	re2 := regexp.MustCompile(`.*COMMAND:`)
	content = re2.ReplaceAllString(content, "")
	content = strings.TrimSpace(content)
	if content == "" {
		panic("empty content")
	}

	storyStep++
	return content
}

func chat(_zm *ZMachine) {
	zm = _zm
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		panic(" OPENAI_API_KEY env var not set")
	}
	client = openai.NewClient(key)
	zm.input = chatInput
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemMsg,
	})

	for !zm.done {
		zm.InterpretInstruction()
		if zm.output.Len() > 0 {
			if zm.windowId == 0 {
				output += zm.output.String()
			}
			zm.output.Reset()
		}

	}

}

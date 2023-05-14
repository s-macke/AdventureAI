package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"regexp"
	"strings"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

var client *openai.Client
var messages []openai.ChatCompletionMessage
var zm *ZMachine

const systemMsg string = `
You act as a player of an interactive text adventure. The goal is to win the game. 
The user provides the text of the text adventure. He is not a human and just prints the output of the game.

Before the command add your thought in curly brackets {}. This thought is mandatory!

Do only write any comments or explanations in curly brackets {}. 
Do only write one command at a time.

The commands always start with a verb. Do never enter an empty command line.
.`

//Do not respond with anything at all other than the command that the player would enter.

func chatInput() string {
	_, _ = os.Stdout.WriteString(zm.output.String())
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: zm.output.String(),
	})
	zm.output.Reset()

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            openai.GPT4,
			Messages:         messages,
			MaxTokens:        128,
			PresencePenalty:  0,
			FrequencyPenalty: 0,
		},
	)
	//fmt.Printf("ChatCompletion response: %v\n", resp)

	if err != nil {

		fmt.Printf("ChatCompletion error: %v\n", err)
		panic("ChatCompletion error")
	}

	re := regexp.MustCompile(`\r?\n`)
	content := re.ReplaceAllString(resp.Choices[0].Message.Content, " ")
	fmt.Printf(InfoColor, content)
	fmt.Println()

	re2 := regexp.MustCompile(`{.*}`)
	content = re2.ReplaceAllString(content, "")
	content = strings.TrimSpace(content)
	if content == "" {
		panic("empty content")
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
	return content
}

func chat(_zm *ZMachine) {
	zm = _zm
	client = openai.NewClient("put your token here")
	zm.input = chatInput
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemMsg,
	})

	for !zm.done {
		zm.InterpretInstruction()
	}

}

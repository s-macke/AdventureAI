package main

import (
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

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

func LoadStoryFromFile(state *StoryState, name string) {
	filename := name
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &state)
	if err != nil {
		panic(err)
	}
}

func PrintPrices(state *StoryState) {
	totalCompletionTokens := 0
	totalPromptTokens := 0
	for i, step := range state.Steps {
		if step.IsResponse {
			totalCompletionTokens += step.CompletionTokens
			totalPromptTokens += step.PromptTokens

			fmt.Println(i, totalCompletionTokens, totalPromptTokens, float32(totalCompletionTokens)*0.00006+float32(totalPromptTokens)*0.00003)
			/*
				for _, message := range step.Messages {
					if message.Completion != "" {
						println(message.Completion)
					}
				}
			*/
		}
	}

}

func PrintStory(state *StoryState) {
	// open input file
	fi, err := os.Create("output.md")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	_, _ = fmt.Fprintln(fi, "# 9:05 by Adam Cadre")
	_, _ = fmt.Fprintln(fi, "")
	_, _ = fmt.Fprintln(fi, "```")
	_, _ = fmt.Fprintln(fi, state.Prompt)
	_, _ = fmt.Fprintln(fi, "```")

	for _, step := range state.Steps {
		if !step.IsResponse {
			continue
		}
		_, _ = fmt.Fprintln(fi, step.Messages[len(step.Messages)-3].Content)
		_, _ = fmt.Fprintln(fi, "> * **Situation:** "+step.Command.Situation)
		_, _ = fmt.Fprintln(fi, "> * **Thought:** "+step.Command.Thought)
		_, _ = fmt.Fprintln(fi, "> * **Command:** "+step.Command.Command)
		_, _ = fmt.Fprintln(fi, "")
	}

}

func main() {
	var state StoryState
	if len(os.Args) < 2 {
		fmt.Println("Usage: extract <filename>")
		os.Exit(0)
	}
	LoadStoryFromFile(&state, os.Args[1])
	//PrintPrices(&state)
	PrintStory(&state)
}

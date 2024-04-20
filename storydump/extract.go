package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Situation string `json:"situation"`
	Thought   string `json:"thought"`
	Command   string `json:"command"`
}

type StoryMessage struct {
	Role             string `json:"role"`
	Content          string `json:"content"`
	CompletionTokens int    `json:"CompletionTokens"`
	PromptTokens     int    `json:"PromptTokens"`
	Meta             string `json:"meta"`
}

type StoryHistory struct {
	PromptPattern string
	Model         string
	Prompt        string
	Messages      []StoryMessage `json:"steps"`
}

func LoadStoryFromFile(state *StoryHistory, name string) {
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

func PrintStory(state *StoryHistory) {
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
	_, _ = fmt.Fprintln(fi, "* played with "+state.Model)
	_, _ = fmt.Fprintln(fi, "```")
	_, _ = fmt.Fprintln(fi, state.Prompt)
	_, _ = fmt.Fprintln(fi, "```")

	for _, message := range state.Messages {
		if message.Role == "user" {
			_, _ = fmt.Fprintln(fi, message.Content)
		}
		if message.Role == "assistant" {
			text := "> * " + message.Meta
			text = strings.Replace(text, "\n\n", "\n", -1)
			text = strings.Replace(text, "\n", "\n> * ", -1)

			_, _ = fmt.Fprintln(fi, text)
			_, _ = fmt.Fprintln(fi, "")
		}
		//_, _ = fmt.Fprintln(fi, step.Messages[len(step.Messages)-3].Content)
		//_, _ = fmt.Fprintln(fi, "> * **Situation:** "+step.Command.Situation)
		//_, _ = fmt.Fprintln(fi, "> * **Thought:** "+step.Command.Thought)
		//_, _ = fmt.Fprintln(fi, "> * **Command:** "+step.Command.Command)
		//_, _ = fmt.Fprintln(fi, "")
	}

}

func main() {
	var state StoryHistory
	if len(os.Args) < 2 {
		fmt.Println("Usage: extract <filename>")
		os.Exit(0)
	}
	LoadStoryFromFile(&state, os.Args[1])
	//PrintPrices(&state)
	PrintStory(&state)
}

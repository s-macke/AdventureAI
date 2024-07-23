package main

import (
	"encoding/json"
	"os"
)

type StoryMessage struct {
	Role             string   `json:"role"`
	Content          string   `json:"content"`
	CompletionTokens int      `json:"completionTokens"`
	PromptTokens     int      `json:"promptTokens"`
	Meta             string   `json:"meta"`
	Score            *float64 `json:"score"`
}

type StoryHistory struct {
	PromptPattern string         `json:"promptPattern"`
	Model         string         `json:"model"`
	Prompt        string         `json:"prompt"`
	Date          string         `json:"date"`
	Name          string         `json:"name"`
	Messages      []StoryMessage `json:"messages"`
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

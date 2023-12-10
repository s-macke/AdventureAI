package chat

import (
	"encoding/json"
	"os"
)

type Command struct {
	Situation string `json:"situation"`
	Narrator  string `json:"narrator"`
	Thought   string `json:"thought"`
	Command   string `json:"command"`
}

type StoryMessage struct {
	Role             string  `json:"role"`
	Content          string  `json:"content"`
	CompletionTokens int     `json:"CompletionTokens"`
	PromptTokens     int     `json:"PromptTokens"`
	IsResponse       bool    `json:"isResponse"`
	Command          Command `json:"command"`
}

type StoryState struct {
	Prompt   string         `json:"prompt"`
	Messages []StoryMessage `json:"steps"`
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
	filename := "storydump/" + name + ".json.backup11"
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &state)
	if err != nil {
		panic(err)
	}
}

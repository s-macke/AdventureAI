package storyHistory

import (
	"encoding/json"
	"os"
)

func (sh *StoryHistory) StoreToFile(name string) {
	stateAsJson, err := json.MarshalIndent(sh, "", " ")
	if err != nil {
		panic(err)
	}
	filename := "storydump/" + name + ".json"
	err = os.WriteFile(filename, stateAsJson, 0644)
	if err != nil {
		panic(err)
	}
}

func (sh *StoryHistory) LoadFromFile(name string) {
	filename := "storydump/" + name + ".json.backup"
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &sh)
	if err != nil {
		panic(err)
	}
}

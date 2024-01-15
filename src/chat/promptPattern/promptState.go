package promptPattern

import (
	"github.com/s-macke/AdventureAI/src/chat/storyHistory"
)

const (
	InfoColor = "\033[1;34m%s\033[0m\n"
)

type State interface {
	GetNextCommand(story *storyHistory.StoryHistory) (string, string)
}

func NewPrompt(stateAsString string, backendAsString string) State {
	switch stateAsString {
	case "simple":
		return NewPromptSimple(backendAsString)
	case "react":
		return NewPromptReAct(backendAsString)
	default:
		panic("Unknown prompt")
	}
}

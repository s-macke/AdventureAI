package promptPattern

import (
	"fmt"
	"github.com/s-macke/AdventureAI/src/chat/storyHistory"
	"strings"
)

const (
	InfoColor = "\033[1;34m%s\033[0m\n"
)

type State interface {
	GetNextCommand(story *storyHistory.StoryHistory) string
}

func NewPrompt(stateAsString string, backendAsString string) State {
	switch stateAsString {
	case "simple":
		return NewPromptSimple(backendAsString)
	case "react":
		return NewPromptReAct(backendAsString)
	case "discuss":
		return NewPromptDiscussion(backendAsString)
	case "history_react":
		return NewPromptHistoryAugmentedReact(backendAsString)
	default:
		panic("Unknown prompt")
	}
}

func CheckAndShowContent(content *string) {
	if *content == "" {
		panic("empty content")
	}
	*content = strings.ReplaceAll(*content, "\r\n", "\n")
	fmt.Printf(InfoColor, *content)
}

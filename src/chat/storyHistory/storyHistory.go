package storyHistory

type StoryMessage struct {
	Role             string `json:"role"`
	Content          string `json:"content"`
	CompletionTokens int    `json:"CompletionTokens"`
	PromptTokens     int    `json:"PromptTokens"`
	Meta             string `json:"meta"`
}

type StoryHistory struct {
	PromptPattern string
	Messages      []StoryMessage `json:"steps"`
}

func (sh *StoryHistory) GetLastMessage() StoryMessage {
	return sh.Messages[len(sh.Messages)-1]
}

func (sh *StoryHistory) AppendMessage(m StoryMessage) {
	sh.Messages = append(sh.Messages, m)
}

func (sh *StoryHistory) GetStory() string {
	var output string
	for _, msg := range sh.Messages {
		output += msg.Content + "\n"
	}
	return output
}

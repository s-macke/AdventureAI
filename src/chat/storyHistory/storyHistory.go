package storyHistory

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

func (sh *StoryHistory) GetLastMessage() StoryMessage {
	return sh.Messages[len(sh.Messages)-1]
}

func (sh *StoryHistory) AppendMessage(m StoryMessage) {
	sh.Messages = append(sh.Messages, m)
}

func (sh *StoryHistory) GetStory() string {
	var output string
	for _, msg := range sh.Messages {
		output += msg.Content
		if msg.Role == "user" {
			output += " "
		}
		if msg.Role == "assistant" {
			output += "\n"
		}
	}
	return output
}

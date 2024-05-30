package backend

const ChatHistoryRoleUser = "user"
const ChatHistoryRoleAssistant = "assistant"

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatHistory struct {
	Messages []ChatMessage `json:"messages"`
}

func (ch *ChatHistory) GetLastMessage() ChatMessage {
	return ch.Messages[len(ch.Messages)-1]
}

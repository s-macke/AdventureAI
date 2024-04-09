package backend

type ChatBackend interface {
	GetResponse(ch *ChatHistory) (string, int, int)
}

func NewChatBackend(prompt string, backendAsString string) ChatBackend {
	switch backendAsString {
	case "gpt3", "gpt4":
		return NewOpenAIChat(prompt, backendAsString)
	case "orca2":
		return NewLlamaChat(prompt, backendAsString)
	case "mistral":
		return NewMistralChat(prompt)
	case "gemini":
		return NewVertexAIChat(prompt)
	case "claude":
		return NewAnthropicChat(prompt)
	case "llama":
		return NewGroqChat(prompt, backendAsString)
	case "gemma":
		return NewGroqChat(prompt, backendAsString)
	default:
		panic("Unknown backend")
	}
}

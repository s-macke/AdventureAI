package backend

type ChatBackend interface {
	GetResponse(ch *ChatHistory) (string, int, int)
}

func NewChatBackend(prompt string, backendAsString string) ChatBackend {
	switch backendAsString {
	case "gpt3", "gpt4", "gpt4o":
		return NewOpenAIChat(prompt, backendAsString)
	case "orca2":
		return NewLlamaChat(prompt, backendAsString)
	case "mistral":
		return NewMistralChat(prompt)
	case "gemini15pro":
		return NewVertexAIChat(prompt, backendAsString)
	case "gemini15flash":
		return NewVertexAIChat(prompt, backendAsString)
	case "opus3":
		return NewAnthropicChat(prompt, backendAsString)
	case "sonnet35":
		return NewAnthropicChat(prompt, backendAsString)
	case "llama":
		return NewLlamaChat(prompt, backendAsString)
	case "gemma":
		return NewGroqChat(prompt, backendAsString)
	default:
		panic("Unknown backend")
	}
}

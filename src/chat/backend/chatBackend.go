package backend

type ChatBackend interface {
	GetResponse(ch *ChatHistory) (string, int, int)
}

func NewChatBackend(prompt string, backendAsString string) ChatBackend {
	switch backendAsString {
	case "gpt-3.5", "gpt-4", "gpt-4-turbo", "gpt-4o", "gpt-4o-mini":
		return NewOpenAIChat(prompt, backendAsString)
	case "orca2":
		return NewLlamaChat(prompt, backendAsString)
	case "mistral-large-2":
		return NewMistralChat(prompt, backendAsString)
	case "gemini-15-pro", "gemini-15-flash", "gemini-15-pro-exp":
		return NewVertexAIChat(prompt, backendAsString)
	case "opus-3", "sonnet-35":
		return NewAnthropicChat(prompt, backendAsString)
	case "llama":
		return NewLlamaChat(prompt, backendAsString)
	case "gemma2":
		return NewGroqChat(prompt, backendAsString)
	case "llama3-8b", "llama3-70b":
		return NewGroqChat(prompt, backendAsString)
	//	case "llama3.1-8b", "llama3.1-70b", "llama3.1-405b":
	//		return NewTogetherChat(prompt, backendAsString)
	//	return NewDeepInfraChat(prompt, backendAsString)
	case "llama3.1-8b", "llama3.1-70b", "llama3.1-405b":
		return NewHyperbolicChat(prompt, backendAsString)
	//case "llama3.1-8b", "llama3.1-70b", "llama3.1-405b":
	//	return NewGroqChat(prompt, backendAsString)
	case "qwen2-72b", "phi3-medium", "phi3-mini":
		return NewDeepInfraChat(prompt, backendAsString)
	default:
		panic("Unknown backend")
	}
}

package backend

type ChatBackend interface {
	GetResponse(ch *ChatHistory) (string, int, int)
}

func NewChatBackend(prompt string, backendAsString string) ChatBackend {
	switch backendAsString {
	case "o3", "o4-mini", "gpt-3.5", "gpt-4", "gpt-4-turbo", "gpt-4o", "gpt-4o-mini":
		return NewOpenAIChat(prompt, backendAsString)
	case "orca2":
		return NewLlamaChat(prompt, backendAsString)
	case "grok-beta":
		return NewXaiChat(prompt, backendAsString)
	case "mistral-large-2":
		return NewMistralChat(prompt, backendAsString)
	case "gemini-2.5-pro", "gemini-2.5-flash", "gemini-2.5-flash-lite":
		return NewVertexAIChat(prompt, backendAsString)
	case "opus-3", "sonnet-35":
		return NewAnthropicChat(prompt, backendAsString)
	case "llama":
		return NewLlamaChat(prompt, backendAsString)
	case "gemma2":
		return NewGroqChat(prompt, backendAsString)
	case "gemma3", "qwen3-0.6b":
		return NewOllamaChat(prompt, backendAsString)
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

package backend

type ChatBackend interface {
	GetResponse(input string) (string, int, int)
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
	default:
		panic("Unknown backend")
	}

}
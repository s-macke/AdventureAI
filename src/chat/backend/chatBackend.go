package backend

import "strings"

type ChatBackend interface {
	GetResponse(ch *ChatHistory) (string, int, int)
}

func NewChatBackend(prompt string, backendAsString string) ChatBackend {
	split := strings.Split(backendAsString, ":")
	if len(split) <= 1 {
		panic("Unknown backend " + backendAsString)
	}
	switch split[0] {
	case "openai":
		return NewOpenAIChat(prompt, split[1])
	case "llama":
		return NewLlamaChat(prompt, split[1])
	case "xai":
		return NewXaiChat(prompt, split[1])
	case "mistral":
		return NewMistralChat(prompt, split[1])
	case "gemini":
		return NewGeminiAIChat(prompt, split[1])
	case "claude":
		return NewAnthropicChat(prompt, split[1])
	case "groq":
		return NewGroqChat(prompt, split[1])
	case "ollama":
		return NewOllamaChat(prompt, split[1])
	case "together":
		return NewTogetherChat(prompt, split[1])
	case "hyperbolic":
		return NewHyperbolicChat(prompt, split[1])
	case "deepinfra":
		return NewDeepInfraChat(prompt, split[1])
	default:
		panic("Unknown backend")
	}
}

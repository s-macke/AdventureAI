package backend

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

type OpenAIChat struct {
	client                *openai.Client
	messages              []openai.ChatCompletionMessage
	totalCompletionTokens int
	totalPromptTokens     int
	CompletionTokens      int
	PromptTokens          int
	model                 string
}

func NewOpenAIChat(systemMsg string, backend string) *OpenAIChat {
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		panic(" OPENAI_API_KEY env var not set")
	}

	cs := &OpenAIChat{
		client: openai.NewClient(key),
	}
	switch backend {
	case "gpt3":
		cs.model = openai.GPT3Dot5Turbo
	case "gpt4":
		cs.model = openai.GPT4TurboPreview
	default:
		panic("Unknown backend")
	}

	cs.messages = append(cs.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemMsg,
	})
	return cs
}

func (cs *OpenAIChat) GetResponse(input string) (string, int, int) {
	cs.messages = append(cs.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: input,
	})
	cs.messages = append(cs.messages, openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleAssistant,
	})

	resp, err := cs.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            cs.model,
			Messages:         cs.messages,
			MaxTokens:        512,
			PresencePenalty:  0,
			FrequencyPenalty: 0,
		},
	)
	cs.totalCompletionTokens += resp.Usage.CompletionTokens
	cs.totalPromptTokens += resp.Usage.PromptTokens
	fmt.Printf("PromptTokens: %d CompletionTokens: %d\n", resp.Usage.PromptTokens, resp.Usage.CompletionTokens)
	fmt.Printf("totalPromptTokens: %d totalCompletionTokens: %d price: %.2f\n",
		cs.totalPromptTokens,
		cs.totalCompletionTokens,
		float32(cs.totalCompletionTokens)*0.00006+float32(cs.totalPromptTokens)*0.00003)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		panic("ChatCompletion error")
	}
	content := resp.Choices[0].Message.Content

	cs.messages = append(cs.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
	return content, resp.Usage.PromptTokens, resp.Usage.CompletionTokens
}

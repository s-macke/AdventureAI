package backend

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"time"
)

type OpenRouterChat struct {
	client    *openai.Client
	model     string
	systemMsg string
}

func NewOpenRouterChat(systemMsg string, backend string) *OpenRouterChat {
	key := os.Getenv("OPENROUTER_API_KEY")
	if key == "" {
		panic("OPENROUTER_API_KEY env var not set")
	}

	config := openai.DefaultConfig(key)
	config.BaseURL = "https://openrouter.ai/api/v1"
	cs := &OpenRouterChat{
		client:    openai.NewClientWithConfig(config),
		systemMsg: systemMsg,
	}

	switch backend {
	case "llama3.1-405b":
		cs.model = "meta-llama/Meta-Llama-3.1-405B-Instruct"
	default:
		panic("Unknown backend")
	}

	return cs
}

func (cs *OpenRouterChat) GetResponse(ch *ChatHistory) (string, int, int) {
	var messages []openai.ChatCompletionMessage
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: cs.systemMsg,
	})

	for _, m := range ch.Messages {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    MapOpenAIRole(m.Role),
			Content: m.Content,
		})
	}

	var resp openai.ChatCompletionResponse
	var err error
	for i := 0; i < 20; i++ {
		resp, err = cs.client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:            cs.model,
				Messages:         messages,
				MaxTokens:        256,
				PresencePenalty:  0,
				FrequencyPenalty: 0,
			},
		)
		if err == nil {
			break
		}
		fmt.Println("ChatCompletion error:", err)
		fmt.Println("Retrying...")
		time.Sleep(20 * time.Second)
	}
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		panic("ChatCompletion error")
	}

	content := resp.Choices[0].Message.Content
	return content, resp.Usage.PromptTokens, resp.Usage.CompletionTokens
}

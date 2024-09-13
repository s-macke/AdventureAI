package backend

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"time"
)

type HyperbolicChat struct {
	client    *openai.Client
	model     string
	systemMsg string
}

func NewHyperbolicChat(systemMsg string, backend string) *HyperbolicChat {
	key := os.Getenv("HYPERBOLIC_API_KEY")
	if key == "" {
		panic("HYPERBOLIC_API_KEY env var not set")
	}

	config := openai.DefaultConfig(key)
	config.BaseURL = "https://api.hyperbolic.xyz/v1"
	cs := &HyperbolicChat{
		client:    openai.NewClientWithConfig(config),
		systemMsg: systemMsg,
	}

	switch backend {
	case "llama3.1-405b":
		cs.model = "meta-llama/Meta-Llama-3.1-405B-Instruct"
	case "llama3.1-70b":
		cs.model = "meta-llama/Meta-Llama-3.1-70B-Instruct"
	case "llama3.1-8b":
		cs.model = "meta-llama/Meta-Llama-3.1-8B-Instruct"
	default:
		panic("Unknown backend")
	}

	return cs
}

func (cs *HyperbolicChat) GetResponse(ch *ChatHistory) (string, int, int) {
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

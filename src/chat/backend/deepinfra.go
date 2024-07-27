package backend

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"time"
)

type DeepInfraChat struct {
	client    *openai.Client
	model     string
	systemMsg string
}

func NewDeepInfraChat(systemMsg string, backend string) *DeepInfraChat {
	key := os.Getenv("DEEPINFRA_TOKEN")
	if key == "" {
		panic("DEEPINFRA_TOKEN env var not set")
	}

	config := openai.DefaultConfig(key)
	config.BaseURL = "https://api.deepinfra.com/v1/openai"
	cs := &DeepInfraChat{
		client:    openai.NewClientWithConfig(config),
		systemMsg: systemMsg,
	}

	switch backend {
	case "llama3.1-8b":
		cs.model = "meta-llama/Meta-Llama-3.1-8B-Instruct"
	case "llama3.1-70b":
		cs.model = "meta-llama/Meta-Llama-3.1-70B-Instruct"
	case "llama3.1-405b":
		cs.model = "meta-llama/Meta-Llama-3.1-405B-Instruct"
	case "qwen2-72b":
		cs.model = "Qwen/Qwen2-72B-Instruct"
	case "phi3-medium":
		cs.model = "microsoft/Phi-3-medium-4k-instruct"
	case "phi3-mini":
		cs.model = "microsoft/Phi-3-mini-4k-instruct"

	default:
		panic("Unknown backend")
	}

	return cs
}

func (cs *DeepInfraChat) GetResponse(ch *ChatHistory) (string, int, int) {
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
				MaxTokens:        1024,
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

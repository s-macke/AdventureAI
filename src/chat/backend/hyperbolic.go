package backend

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type HyperbolicChat struct {
	client    openai.Client
	model     string
	systemMsg string
}

func NewHyperbolicChat(systemMsg string, backend string) *HyperbolicChat {
	key := os.Getenv("HYPERBOLIC_API_KEY")
	if key == "" {
		panic("HYPERBOLIC_API_KEY env var not set")
	}

	cs := &HyperbolicChat{
		client: openai.NewClient(
			option.WithAPIKey(key),
			option.WithBaseURL("https://api.hyperbolic.xyz/v1"),
		),
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
	var messages []openai.ChatCompletionMessageParamUnion
	messages = append(messages, openai.SystemMessage(cs.systemMsg))

	for _, m := range ch.Messages {
		messages = append(messages, MapOpenAIRole(m.Role, m.Content))
	}

	var resp *openai.ChatCompletion
	var err error
	for i := 0; i < 20; i++ {
		resp, err = cs.client.Chat.Completions.New(
			context.Background(),
			openai.ChatCompletionNewParams{
				Model:     openai.ChatModel(cs.model),
				Messages:  messages,
				MaxTokens: openai.Int(256),
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
	return content, int(resp.Usage.PromptTokens), int(resp.Usage.CompletionTokens)
}

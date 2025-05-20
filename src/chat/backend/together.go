package backend

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type TogetherChat struct {
	client    openai.Client
	model     string
	systemMsg string
}

func NewTogetherChat(systemMsg string, backend string) *TogetherChat {
	key := os.Getenv("TOGETHER_API_KEY")
	if key == "" {
		panic("TOGETHER_API_KEY env var not set")
	}

	cs := &TogetherChat{
		client: openai.NewClient(
			option.WithAPIKey(key),
			option.WithBaseURL("https://api.together.xyz/v1"),
		),
		systemMsg: systemMsg,
	}

	switch backend {
	case "llama3.1-8b":
		cs.model = "meta-llama/Meta-Llama-3.1-8B-Instruct-Turbo"
	case "llama3.1-70b":
		cs.model = "meta-llama/Meta-Llama-3.1-70B-Instruct-Turbo"
	case "llama3.1-405b":
		cs.model = "meta-llama/Meta-Llama-3.1-405B-Instruct-Turbo"
	default:
		panic("Unknown backend")
	}

	return cs
}

func (cs *TogetherChat) GetResponse(ch *ChatHistory) (string, int, int) {
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
				MaxTokens: openai.Int(1024),
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

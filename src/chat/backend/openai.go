package backend

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"time"
)

type OpenAIChat struct {
	client                *openai.Client
	totalCompletionTokens int
	totalPromptTokens     int
	CompletionTokens      int
	PromptTokens          int
	model                 string
	systemMsg             string
}

func NewOpenAIChat(systemMsg string, backend string) *OpenAIChat {
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		panic(" OPENAI_API_KEY env var not set")
	}

	cs := &OpenAIChat{
		client:    openai.NewClient(key),
		systemMsg: systemMsg,
	}
	switch backend {
	case "gpt-3.5":
		cs.model = openai.GPT3Dot5Turbo
	case "gpt-4-turbo":
		cs.model = openai.GPT4Turbo
	case "gpt-4":
		cs.model = openai.GPT4
	case "gpt-4o":
		cs.model = openai.GPT4o
	case "gpt-4o-mini":
		cs.model = openai.GPT4oMini
	default:
		panic("Unknown backend")
	}

	return cs
}

func (cs *OpenAIChat) GetResponse(ch *ChatHistory) (string, int, int) {
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
	for i := 0; i < 10; i++ {
		resp, err = cs.client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:            cs.model,
				Messages:         messages,
				MaxTokens:        2048,
				PresencePenalty:  0,
				FrequencyPenalty: 0,
			},
		)
		if err == nil {
			break
		}
		fmt.Println("ChatCompletion error:", err)
		fmt.Println("Retrying...")
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		panic("ChatCompletion error")
	}

	/*
		d := time.Since(start)
		f, err := os.OpenFile("text2.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		_, _ = f.WriteString(fmt.Sprintf("%d %d %.2f\n",
			resp.Usage.CompletionTokens,
			resp.Usage.PromptTokens,
			float64(d)/1.e9))
	*/
	content := resp.Choices[0].Message.Content
	return content, resp.Usage.PromptTokens, resp.Usage.CompletionTokens
}

func MapOpenAIRole(role string) string {
	switch role {
	case ChatHistoryRoleUser:
		return openai.ChatMessageRoleUser
	case ChatHistoryRoleAssistant:
		return openai.ChatMessageRoleAssistant
	default:
		panic("Unknown role")
	}
}

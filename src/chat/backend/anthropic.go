package backend

import (
	"context"
	"fmt"
	"github.com/liushuangls/go-anthropic/v2"
	"os"
	"time"
)

type AnthropicChat struct {
	client *anthropic.Client
	prompt string
	model  string
}

func NewAnthropicChat(systemMsg string, backend string) *AnthropicChat {
	key := os.Getenv("ANTHROPIC_API_KEY")
	if key == "" {
		panic("ANTHROPIC_API_KEY env var not set")
	}
	c := anthropic.NewClient(key)
	return &AnthropicChat{
		client: c,
		prompt: systemMsg,
		model:  backend,
	}
}

func (cs *AnthropicChat) GetResponse(ch *ChatHistory) (string, int, int) {
	var messages []anthropic.Message
	for _, m := range ch.Messages {
		messages = append(messages, anthropic.Message{
			Role:    MapAnthropicRole(m.Role),
			Content: []anthropic.MessageContent{anthropic.NewTextMessageContent(m.Content)},
		})
	}
	var response anthropic.MessagesResponse
	var err error
	// retries
	for i := 0; i < 20; i++ {
		request := anthropic.MessagesRequest{
			Messages:  messages,
			MaxTokens: 4096,
			System:    cs.prompt,
		}
		switch cs.model {
		case "opus-3":
			request.Model = "claude-3-opus-20240229"
		case "sonnet-35":
			request.Model = "claude-3-5-sonnet-20240620"
		case "sonnet-4":
			request.Model = "claude-sonnet-4-20250514"
		case "opus-4":
			request.Model = "claude-opus-4-20250514"
		default:
			panic("Unknown model")
		}

		response, err = cs.client.CreateMessages(
			context.Background(), request)
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(15 * time.Second)
	}
	if err != nil {
		panic(err)
	}

	return response.GetFirstContentText(), response.Usage.InputTokens, response.Usage.OutputTokens
}

func MapAnthropicRole(role string) anthropic.ChatRole {
	switch role {
	case ChatHistoryRoleUser:
		return anthropic.RoleUser
	case ChatHistoryRoleAssistant:
		return anthropic.RoleAssistant
	default:
		panic("Unknown role")
	}
}

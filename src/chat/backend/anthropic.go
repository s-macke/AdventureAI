package backend

import (
	"context"
	"fmt"
	"github.com/liushuangls/go-anthropic"
	"os"
	"time"
)

type AnthropicChat struct {
	client *anthropic.Client
	prompt string
}

func NewAnthropicChat(systemMsg string) *AnthropicChat {
	key := os.Getenv("ANTHROPIC_API_KEY")
	if key == "" {
		panic("ANTHROPIC_API_KEY env var not set")
	}
	c := anthropic.NewClient(key)
	return &AnthropicChat{
		client: c,
		prompt: systemMsg,
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
	for i := 0; i < 5; i++ {
		response, err = cs.client.CreateMessages(
			context.Background(),
			anthropic.MessagesRequest{
				//Model:     "claude-3-opus-20240229",
				Model:     "claude-3-5-sonnet-20240620",
				Messages:  messages,
				MaxTokens: 4096,
				System:    cs.prompt,
			})
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		panic(err)
	}

	return response.GetFirstContentText(), response.Usage.InputTokens, response.Usage.OutputTokens
}

func MapAnthropicRole(role string) string {
	switch role {
	case ChatHistoryRoleUser:
		return anthropic.RoleUser
	case ChatHistoryRoleAssistant:
		return anthropic.RoleAssistant
	default:
		panic("Unknown role")
	}
}

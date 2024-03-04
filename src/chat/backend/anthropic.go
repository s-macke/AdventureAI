package backend

import (
	"context"
	"github.com/liushuangls/go-anthropic"
	"os"
)

type AnthropicChat struct {
	client   *anthropic.Client
	messages []anthropic.Message
	prompt   string
}

func NewAnthropicChat(systemMsg string) *AnthropicChat {
	key := os.Getenv("ANTHROPIC_API_KEY")
	if key == "" {
		panic("ANTHROPIC_API_KEY env var not set")
	}
	c := anthropic.NewClient(key)
	return &AnthropicChat{
		client:   c,
		messages: make([]anthropic.Message, 0),
		prompt:   systemMsg,
	}
}

func (cs *AnthropicChat) GetResponse(input string) (string, int, int) {
	cs.messages = append(cs.messages, anthropic.Message{
		Role:    anthropic.RoleUser,
		Content: input,
	})

	response, err := cs.client.CreateMessages(
		context.Background(),
		anthropic.MessagesRequest{
			Model:     "claude-3-opus-20240229",
			Messages:  cs.messages,
			MaxTokens: 2048,
			System:    cs.prompt,
		})
	if err != nil {
		panic(err)
	}
	cs.messages = append(cs.messages, anthropic.Message{
		Role:    anthropic.RoleAssistant,
		Content: response.Content[0].Text,
	})

	return response.Content[0].Text, response.Usage.InputTokens, response.Usage.OutputTokens

}

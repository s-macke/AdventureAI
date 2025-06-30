package backend

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OllamaChat struct {
	client    openai.Client
	model     openai.ChatModel
	systemMsg string
}

func NewOllamaChat(systemMsg string, backend string) *OllamaChat {
	cs := &OllamaChat{
		client:    openai.NewClient(option.WithBaseURL("http://localhost:11434/v1")),
		systemMsg: systemMsg,
	}
	switch backend {
	case "gemma3":
		cs.model = "gemma3:27b-it-qat"
	case "qwen3-0.6b":
		cs.model = "qwen3:0.6b-fp16"

	default:
		panic("Unknown backend")
	}

	return cs
}

func (cs *OllamaChat) GetResponse(ch *ChatHistory) (string, int, int) {

	var messages []openai.ChatCompletionMessageParamUnion
	messages = append(messages,
		openai.SystemMessage(cs.systemMsg),
	)

	for _, m := range ch.Messages {
		switch m.Role {
		case ChatHistoryRoleUser:
			messages = append(messages, openai.UserMessage(m.Content))

		case ChatHistoryRoleAssistant:
			messages = append(messages, openai.AssistantMessage(m.Content))

		default:
			panic("Unknown role")
		}
	}
	var chatCompletion *openai.ChatCompletion
	var err error
	for i := 0; i < 3; i++ {
		chatCompletion, err = cs.client.Chat.Completions.New(context.TODO(),
			openai.ChatCompletionNewParams{
				Messages: messages,
				Model:    cs.model,
			})
		if err != nil {
			fmt.Println("ChatCompletion error:", err)
			fmt.Println("Retrying...")
			continue
		}
		if chatCompletion == nil || (*chatCompletion).Choices[0].Message.Content == "" {
			fmt.Println("ChatCompletion error: No output")
			continue
		}
		break
	}
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		panic("ChatCompletion error")
	}
	return (*chatCompletion).Choices[0].Message.Content, 0, 0
}

package backend

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

type OpenAIChat struct {
	client                *openai.Client
	messages              []openai.ChatCompletionMessage
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
	case "gpt3":
		cs.model = openai.GPT3Dot5Turbo
	case "gpt4":
		cs.model = openai.GPT4TurboPreview
	default:
		panic("Unknown backend")
	}

	cs.messages = append(cs.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemMsg,
	})
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
	//start := time.Now()
	resp, err := cs.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            cs.model,
			Messages:         messages,
			MaxTokens:        2048,
			PresencePenalty:  0,
			FrequencyPenalty: 0,
		},
	)
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
	case "user":
		return openai.ChatMessageRoleUser
	case "assistant":
		return openai.ChatMessageRoleAssistant
	default:
		panic("Unknown role")
	}
}

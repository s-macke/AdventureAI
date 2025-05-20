package backend

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIChat struct {
	client                openai.Client
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
		client:    openai.NewClient(option.WithAPIKey(key)),
		systemMsg: systemMsg,
	}
	switch backend {
	case "gpt-3.5":
		cs.model = string(openai.ChatModelGPT3_5Turbo)
	case "o1-preview":
		cs.model = string(openai.ChatModelO1Preview)
	case "o1":
		cs.model = string(openai.ChatModelO1)
	case "o3-mini":
		cs.model = string(openai.ChatModelO3Mini)
	case "o1-mini":
		cs.model = string(openai.ChatModelO1Mini)
	case "gpt-4-turbo":
		cs.model = string(openai.ChatModelGPT4Turbo)
	case "gpt-4":
		cs.model = string(openai.ChatModelGPT4)
	case "gpt-4o":
		cs.model = string(openai.ChatModelGPT4o)
	case "gpt-4o-mini":
		cs.model = string(openai.ChatModelGPT4oMini)
	case "o3":
		cs.model = "o3"
	case "o4-mini":
		cs.model = "o4-mini"
	default:
		panic("Unknown backend")
	}

	return cs
}

func (cs *OpenAIChat) GetResponse(ch *ChatHistory) (string, int, int) {
	var messages []openai.ChatCompletionMessageParamUnion
	messages = append(messages, openai.SystemMessage(cs.systemMsg))

	for _, m := range ch.Messages {
		messages = append(messages, MapOpenAIRole(m.Role, m.Content))
	}

	var resp *openai.ChatCompletion
	var err error
	for i := 0; i < 10; i++ {
		resp, err = cs.client.Chat.Completions.New(
			context.Background(),
			openai.ChatCompletionNewParams{
				Model:    openai.ChatModel(cs.model),
				Messages: messages,
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
	return content, int(resp.Usage.PromptTokens), int(resp.Usage.CompletionTokens)
}

func MapOpenAIRole(role, content string) openai.ChatCompletionMessageParamUnion {
	switch role {
	case ChatHistoryRoleUser:
		return openai.UserMessage(content)
	case ChatHistoryRoleAssistant:
		return openai.AssistantMessage(content)
	default:
		panic("Unknown role")
	}
}

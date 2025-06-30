package backend

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/responses"
	"os"
	"time"
)

type OpenAIChat struct {
	client                openai.Client
	model                 responses.ResponsesModel
	totalCompletionTokens int
	totalPromptTokens     int
	CompletionTokens      int
	PromptTokens          int
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
		cs.model = openai.ChatModelGPT3_5Turbo
	case "o1-preview":
		cs.model = openai.ChatModelO1Preview
	case "o1":
		cs.model = openai.ChatModelO1
	case "o3-mini":
		cs.model = openai.ChatModelO3Mini
	case "o1-mini":
		cs.model = openai.ChatModelO1Mini
	case "gpt-4-turbo":
		cs.model = openai.ChatModelGPT4Turbo
	case "gpt-4":
		cs.model = openai.ChatModelGPT4
	case "gpt-4o":
		cs.model = openai.ChatModelGPT4o
	case "gpt-4o-mini":
		cs.model = openai.ChatModelGPT4oMini
	case "o3":
		cs.model = openai.ChatModelO3
	case "o4-mini":
		cs.model = openai.ChatModelO4Mini
	default:
		panic("Unknown backend")
	}

	return cs
}

func (cs *OpenAIChat) GetResponse(ch *ChatHistory) (string, int, int) {

	var messages responses.ResponseInputParam
	for _, m := range ch.Messages {
		messages = append(messages, responses.ResponseInputItemUnionParam{
			OfMessage: &responses.EasyInputMessageParam{
				Role: MapOpenAIResponsesRole(m.Role),
				Content: responses.EasyInputMessageContentUnionParam{
					OfString: openai.Opt(m.Content),
				},
			},
		})
	}

	var responseCompletion *responses.Response
	var err error
	for i := 0; i < 3; i++ {
		responseCompletion, err = cs.client.Responses.New(
			context.Background(),
			responses.ResponseNewParams{
				Instructions: param.Opt[string]{
					Value: cs.systemMsg,
				},
				Input: responses.ResponseNewParamsInputUnion{
					OfInputItemList: messages,
				},
				Model: cs.model,
			})
		if err != nil {
			fmt.Println("Responses error:", err)
			fmt.Println("Retrying...")
			time.Sleep(2 * time.Second)
			continue
		}
		if responseCompletion == nil || responseCompletion.OutputText() == "" {
			fmt.Println("Responses error: No output")
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	if err != nil || responseCompletion == nil {
		fmt.Printf("Responses error: %v\n", err)
		panic("Responses error")
	}
	return responseCompletion.OutputText(), 0, 0
}

func MapOpenAIResponsesRole(role string) responses.EasyInputMessageRole {
	switch role {
	case ChatHistoryRoleUser:
		return responses.EasyInputMessageRoleUser
	case ChatHistoryRoleAssistant:
		return responses.EasyInputMessageRoleAssistant
	default:
		panic("Unknown role")
	}
}

package backend

import (
	"fmt"
	"github.com/gage-technologies/mistral-go"
	"os"
	"time"
)

type MistralChat struct {
	client *mistral.MistralClient
	prompt string
	model  string
}

func NewMistralChat(systemMsg string, backend string) *MistralChat {
	key := os.Getenv("MISTRAL_API_KEY")
	if key == "" {
		panic(" MISTRAL_API_KEY env var not set")
	}

	cs := &MistralChat{
		client: mistral.NewMistralClientDefault(key),
	}
	switch backend {
	case "mistral-large-2":
		cs.model = "mistral-large-2407"
	default:
		panic("Unknown backend")

	}
	//cs.model = "open-mistral-7b"
	//cs.model = "open-mixtral-8x7b"
	//cs.model = "open-mixtral-8x22b"
	//cs.model = "open-mixtral-8x22b"
	return cs

}

func (cs *MistralChat) GetResponse(ch *ChatHistory) (string, int, int) {
	var messages []mistral.ChatMessage

	messages = append(messages, mistral.ChatMessage{
		Role:    mistral.RoleSystem,
		Content: cs.prompt,
	})

	for _, m := range ch.Messages {
		messages = append(messages, mistral.ChatMessage{
			Role:    MapMistralRole(m.Role),
			Content: m.Content,
		})
	}

	var response *mistral.ChatCompletionResponse
	var err error
	for i := 0; i < 10; i++ {
		response, err = cs.client.Chat(cs.model, messages, &mistral.DefaultChatRequestParams)
		if err == nil {
			break
		}
		fmt.Println("ChatCompletion error:", err)
		fmt.Println("Retrying...")
		time.Sleep(10 * time.Second)
	}
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		panic("ChatCompletion error")
	}
	return response.Choices[0].Message.Content, response.Usage.PromptTokens, response.Usage.TotalTokens

}

func MapMistralRole(role string) string {
	switch role {
	case ChatHistoryRoleUser:
		return mistral.RoleUser
	case ChatHistoryRoleAssistant:
		return mistral.RoleAssistant
	default:
		panic("Unknown role")
	}
}

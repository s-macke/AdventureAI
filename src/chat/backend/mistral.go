package backend

import (
	"github.com/gage-technologies/mistral-go"
	"os"
)

type MistralChat struct {
	messages              []mistral.ChatMessage
	totalCompletionTokens int
	totalPromptTokens     int
	client                *mistral.MistralClient
}

func NewMistralChat(systemMsg string) *MistralChat {
	key := os.Getenv("MISTRAL_API_KEY")
	if key == "" {
		panic(" MISTRAL_API_KEY env var not set")
	}

	cs := &MistralChat{
		client: mistral.NewMistralClientDefault(key),
	}

	cs.messages = append(cs.messages, mistral.ChatMessage{
		Role:    mistral.RoleSystem,
		Content: systemMsg,
	})

	return cs

}

func (cs *MistralChat) GetResponse(input string) (string, int, int) {
	cs.messages = append(cs.messages, mistral.ChatMessage{
		Role:    mistral.RoleUser,
		Content: input,
	})
	/*
		"mistral-medium"
		"mistral-small"
		"mistral-tiny"
	*/
	response, err := cs.client.Chat("mistral-medium", cs.messages, &mistral.DefaultChatRequestParams)
	if err != nil {
		panic(err)
	}

	cs.messages = append(cs.messages, mistral.ChatMessage{
		Role:    mistral.RoleAssistant,
		Content: response.Choices[0].Message.Content,
	})

	return response.Choices[0].Message.Content, response.Usage.PromptTokens, response.Usage.TotalTokens
}

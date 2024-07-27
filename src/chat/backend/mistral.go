package backend

import (
	"github.com/gage-technologies/mistral-go"
	"os"
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
	cs.model = "open-mistral-7b"
	cs.model = "open-mixtral-8x7b"
	cs.model = "open-mixtral-8x22b"
	cs.model = "open-mixtral-8x22b"

	return cs

}

func (cs *MistralChat) GetResponse(ch *ChatHistory) (string, int, int) {
	panic("not implemented")
	/*
		cs.messages = append(cs.messages, mistral.ChatMessage{
			Role:    mistral.RoleUser,
			Content: input,
		})
		/*
				"mistral-medium"
				"mistral-small"
				"mistral-tiny"
			    "mistral-large"
	*/
	/*
		response, err := cs.client.Chat("mistral-medium", cs.messages, &mistral.DefaultChatRequestParams)
		if err != nil {
			panic(err)
		}

		cs.messages = append(cs.messages, mistral.ChatMessage{
			Role:    mistral.RoleAssistant,
			Content: response.Choices[0].Message.Content,
		})

		return response.Choices[0].Message.Content, response.Usage.PromptTokens, response.Usage.TotalTokens
	*/
}

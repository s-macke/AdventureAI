package backend

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"os"
)

type VertexAIChat struct {
	ctx          context.Context
	client       *genai.Client
	gemini       *genai.GenerativeModel
	systemPrompt string
}

func NewVertexAIChat(systemMsg string) *VertexAIChat {
	key := os.Getenv("GEMINI_API_KEY")
	if key == "" {
		panic("GEMINI_API_KEY env var not set")
	}

	var err error
	cs := &VertexAIChat{
		ctx:          context.Background(),
		systemPrompt: systemMsg,
	}
	cs.client, err = genai.NewClient(cs.ctx, option.WithAPIKey(key))
	if err != nil {
		panic(err)
	}
	//cs.gemini = cs.client.GenerativeModel("gemini-pro") // reference to gemini-1.0-pro
	//cs.gemini = cs.client.GenerativeModel("gemini-1.5-pro")
	cs.gemini = cs.client.GenerativeModel("gemini-1.0-pro")
	safety := []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryUnspecified,
			Threshold: genai.HarmBlockOnlyHigh,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockOnlyHigh,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockOnlyHigh,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockOnlyHigh,
		},
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockOnlyHigh,
		},
	}
	cs.gemini.SafetySettings = safety
	return cs
}

func (cs *VertexAIChat) GetResponse(ch *ChatHistory) (string, int, int) {
	var messages []*genai.Content

	if ch.GetLastMessage().Role != "user" {
		panic("Last message must be from user")
	}
	lastMessage := genai.Text(ch.GetLastMessage().Content)
	if len(ch.Messages) == 1 {
		lastMessage = genai.Text(cs.systemPrompt) + "\n\n---\n\n" + lastMessage
	}

	for i, m := range ch.Messages {
		// Don't add last message
		if i == len(ch.Messages)-1 {
			break
		}
		if i == 0 {
			messages = append(messages, &genai.Content{
				Role:  MapVertexAiRole(m.Role),
				Parts: []genai.Part{genai.Text(cs.systemPrompt) + "\n\n---\n\n" + genai.Text(m.Content)},
			})
		} else {
			messages = append(messages, &genai.Content{
				Role:  MapVertexAiRole(m.Role),
				Parts: []genai.Part{genai.Text(m.Content)},
			})
		}
	}

	chat := cs.gemini.StartChat()
	chat.History = messages
	response, err := chat.SendMessage(cs.ctx, lastMessage)
	if err != nil {
		panic(err)
	}
	//fmt.Println(response.Candidates[0].Content.Parts[0].(genai.Text))
	output := response.Candidates[0].Content.Parts[0].(genai.Text)
	return string(output), 0, 0
}

func MapVertexAiRole(role string) string {
	switch role {
	case "user":
		return "user"
	case "assistant":
		return "model"
	default:
		panic("Unknown role")
	}
}

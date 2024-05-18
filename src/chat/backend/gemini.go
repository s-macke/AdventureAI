package backend

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"os"
	"strings"
	"time"
)

type VertexAIChat struct {
	ctx    context.Context
	client *genai.Client
	gemini *genai.GenerativeModel
}

func NewVertexAIChat(systemMsg string) *VertexAIChat {
	key := os.Getenv("GEMINI_API_KEY")
	if key == "" {
		panic("GEMINI_API_KEY env var not set")
	}

	var err error
	cs := &VertexAIChat{
		ctx: context.Background(),
	}
	cs.client, err = genai.NewClient(cs.ctx, option.WithAPIKey(key))
	if err != nil {
		panic(err)
	}
	/*
		it := cs.client.ListModels(context.Background())
		for {
			model, err := it.Next()
			if err == nil {
				fmt.Println(model)
			} else {
				break
			}
		}
	*/
	//cs.gemini = cs.client.GenerativeModel("gemini-pro") // reference to gemini-1.0-pro
	//cs.gemini = cs.client.GenerativeModel("gemini-1.5-flash-latest")
	cs.gemini = cs.client.GenerativeModel("gemini-1.5-pro-latest")
	cs.gemini.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(systemMsg)},
	}

	//cs.gemini = cs.client.GenerativeModel("gemini-1.0-pro")
	safety := []*genai.SafetySetting{
		/*
			{
				Category:  genai.HarmCategoryUnspecified,
				Threshold: genai.HarmBlockMediumAndAbove,
			},
			{
				Category:  genai.HarmCategoryDerogatory,
				Threshold: genai.HarmBlockMediumAndAbove,
			},
			{
				Category:  genai.HarmCategoryToxicity,
				Threshold: genai.HarmBlockMediumAndAbove,
			},
			{
				Category:  genai.HarmCategoryViolence,
				Threshold: genai.HarmBlockMediumAndAbove,
			},
			{
				Category:  genai.HarmCategorySexual,
				Threshold: genai.HarmBlockMediumAndAbove,
			},
			{
				Category:  genai.HarmCategoryMedical,
				Threshold: genai.HarmBlockMediumAndAbove,
			},
			{
				Category:  genai.HarmCategoryDangerous,
				Threshold: genai.HarmBlockMediumAndAbove,
			},
		*/
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockNone,
		},
	}
	cs.gemini.SafetySettings = safety
	return cs
}

func (cs *VertexAIChat) CallChat(messages []*genai.Content, lastMessage genai.Text) *genai.GenerateContentResponse {
	var err error
	var response *genai.GenerateContentResponse
	for i := 0; i < 10; i++ {
		chat := cs.gemini.StartChat()
		chat.History = messages
		response, err = chat.SendMessage(cs.ctx, lastMessage)
		if err != nil {
			fmt.Println("Error:", err)
			if strings.Contains(err.Error(), "429") {
				time.Sleep(20 * time.Second)
				continue
			}
			/*
				if s, ok := status.FromError(err); ok {
					if s.Code() == 429 {
						time.Sleep(20 * time.Second)
						continue
					}
				}
			*/
			panic(errors.Unwrap(err))
		}
		return response
	}
	panic(err)
}

func (cs *VertexAIChat) GetResponse(ch *ChatHistory) (string, int, int) {
	var messages []*genai.Content

	if ch.GetLastMessage().Role != "user" {
		panic("Last message must be from user")
	}

	lastMessage := genai.Text(ch.GetLastMessage().Content)

	for i, m := range ch.Messages {
		// Don't add last message
		if i == len(ch.Messages)-1 {
			break
		}
		messages = append(messages, &genai.Content{
			Role:  MapVertexAiRole(m.Role),
			Parts: []genai.Part{genai.Text(m.Content)},
		})
	}
	response := cs.CallChat(messages, lastMessage)

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

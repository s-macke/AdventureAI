package backend

import (
	"context"
	"google.golang.org/genai"
	"os"
)

type GeminiAIChat struct {
	ctx    context.Context
	client *genai.Client
	config *genai.GenerateContentConfig
	model  string
}

func NewVertexAIChat(systemMsg string, model string) *GeminiAIChat {
	key := os.Getenv("GEMINI_API_KEY")
	if key == "" {
		panic("GEMINI_API_KEY env var not set")
	}

	var err error
	cs := &GeminiAIChat{
		ctx: context.Background(),
	}
	cs.client, err = genai.NewClient(cs.ctx, &genai.ClientConfig{
		APIKey:  key,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		panic(err)
	}
	/*
		for key := range cs.client.Models.All(cs.ctx) {
			fmt.Println(key.Name)
		}
		os.Exit(0)
	*/
	switch model {
	case "gemini-15-pro":
		cs.model = "gemini-1.5-pro-latest"
	case "gemini-15-pro-exp":
		cs.model = "gemini-1.5-pro-exp-0801"
	case "gemini-15-flash":
		cs.model = "gemini-1.5-flash-latest"
	case "gemini-20-flash":
		cs.model = "gemini-2.0-flash"
	default:
		panic("Unknown model: " + model)
	}

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
			Threshold: genai.HarmBlockThresholdBlockNone,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockThresholdBlockNone,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockThresholdBlockNone,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockThresholdBlockNone,
		},
	}

	cs.config = &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Role:  genai.RoleUser,
			Parts: []*genai.Part{{Text: systemMsg}},
		},
		SafetySettings: safety,
	}

	return cs
}

func (cs *GeminiAIChat) GetResponse(ch *ChatHistory) (string, int, int) {
	var history []*genai.Content

	for _, m := range ch.Messages {
		history = append(history, &genai.Content{
			Role:  string(MapGeminiAiRole(m.Role)),
			Parts: []*genai.Part{{Text: m.Content}},
		})
	}
	response, err := cs.client.Models.GenerateContent(cs.ctx, cs.model, history, cs.config)
	if err != nil {
		panic(err)
	}
	return response.Text(), 0, 0
}

func MapGeminiAiRole(role string) genai.Role {
	switch role {
	case ChatHistoryRoleUser:
		return genai.RoleUser
	case ChatHistoryRoleAssistant:
		return genai.RoleModel
	default:
		panic("Unknown role")
	}
}

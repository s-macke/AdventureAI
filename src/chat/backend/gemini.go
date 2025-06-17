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
	case "gemini-2.5-pro":
		cs.model = "gemini-2.5-pro"
	case "gemini-2.5-flash":
		cs.model = "gemini-2.5-flash"
	case "gemini-2.5-flash-lite":
		cs.model = "gemini-2.5-flash-lite-preview-06-17"
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
	//budget := int32(-1) // -1 means no budget limit
	budget := int32(0) // 0 means no thinking
	cs.config = &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Role:  genai.RoleUser,
			Parts: []*genai.Part{{Text: systemMsg}},
		},
		SafetySettings: safety,
		ThinkingConfig: &genai.ThinkingConfig{
			IncludeThoughts: false,
			ThinkingBudget:  &budget,
		},
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

package backend

import (
	"cloud.google.com/go/vertexai/genai"
	"context"
	"os"
)

type VertexAIChat struct {
	client       *genai.Client
	gemini       *genai.GenerativeModel
	chat         *genai.ChatSession
	systemPrompt string
	first        bool
}

func NewVertexAIChat(systemMsg string) *VertexAIChat {
	key := os.Getenv("GOOGLE_PROJECT_ID")
	if key == "" {
		panic("GOOGLE_PROJECT_ID env var not set")
	}

	var err error
	cs := &VertexAIChat{
		systemPrompt: systemMsg,
		first:        true,
	}

	cs.client, err = genai.NewClient(context.TODO(), key, "us-central1")
	if err != nil {
		panic(err)
	}
	cs.gemini = cs.client.GenerativeModel("gemini-pro")
	safety := []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryUnspecified,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockNone,
		},
	}
	cs.gemini.SafetySettings = safety
	cs.chat = cs.gemini.StartChat()

	//parts := []genai.Part{genai.Text(systemMsg)}
	//cs.chat.History = append(cs.chat.History, &genai.Content{Role: "user", Parts: parts})
	//rb, _ := json.MarshalIndent(response, "", "  ")
	//fmt.Println(string(rb))
	return cs
}

func (cs *VertexAIChat) GetResponse(input string) (string, int, int) {
	var text string
	if cs.first {
		cs.first = false
		text = cs.systemPrompt + "\n\n\n---\n\n" + input
	} else {
		text = input
	}
	response, err := cs.chat.SendMessage(context.TODO(), genai.Text(text))
	if err != nil {
		panic(err)
	}
	//fmt.Println(response.Candidates[0].Content.Parts[0].(genai.Text))
	output := response.Candidates[0].Content.Parts[0].(genai.Text)
	return string(output), 0, 0
}

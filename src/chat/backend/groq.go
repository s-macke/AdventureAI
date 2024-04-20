package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type GroqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqRequest struct {
	Messages []GroqMessage `json:"messages"`
	Model    string        `json:"model"`
}

type GroqResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int     `json:"prompt_tokens"`
		PromptTime       float64 `json:"prompt_time"`
		CompletionTokens int     `json:"completion_tokens"`
		CompletionTime   float64 `json:"completion_time"`
		TotalTokens      int     `json:"total_tokens"`
		TotalTime        float64 `json:"total_time"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
	XGroq             struct {
		ID string `json:"id"`
	} `json:"x_groq"`
}

type GroqChat struct {
	apikey string
	prompt string
	model  string
}

func NewGroqChat(systemMsg string, backend string) *GroqChat {
	key := os.Getenv("GROQ_API_KEY")
	if key == "" {
		panic(" GROQ_API_KEY env var not set")
	}

	cs := &GroqChat{
		apikey: key,
		prompt: systemMsg,
	}
	switch backend {
	case "llama":
		cs.model = "llama3-8b-8192"
	//case "llama":	cs.model = "llama3-70b-8192"
	case "gemma":
		cs.model = "gemma-7b-it"
	}

	return cs
}

func (cs *GroqChat) CallGroq(request *GroqRequest) (*GroqResponse, error) {
	data, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cs.apikey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusTooManyRequests {
			fmt.Println("Rate limit exceeded. Waiting 30 seconds")
			time.Sleep(30 * time.Second)
			return nil, fmt.Errorf("Rate limit exceeded")
		}
		panic(res.Status)
	}

	response := &GroqResponse{}
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

func (cs *GroqChat) GetResponse(ch *ChatHistory) (string, int, int) {
	request := GroqRequest{
		Messages: []GroqMessage{},
		Model:    cs.model,
	}

	request.Messages = append(request.Messages, GroqMessage{
		Role:    "system",
		Content: cs.prompt,
	})

	for _, m := range ch.Messages {
		request.Messages = append(request.Messages, GroqMessage{
			Role:    m.Role,
			Content: m.Content,
		})
	}
	for attempts := 0; attempts < 10; attempts++ {
		response, err := cs.CallGroq(&request)
		if err == nil {
			return response.Choices[0].Message.Content, 0, 0
		}
	}
	panic("Too many attempts. failed to get response")
}

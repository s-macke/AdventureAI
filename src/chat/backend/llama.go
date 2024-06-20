package backend

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type LlamaChat struct {
	Prompt string
}

type LlamaRequest struct {
	CachePrompt bool     `json:"cache_prompt"`
	NPredict    int      `json:"n_predict"`
	Prompt      string   `json:"prompt"`
	Stop        []string `json:"stop"`
	Stream      bool     `json:"stream"`
}

/*
type LlamaRequest struct {
	CachePrompt      bool     `json:"cache_prompt"`
	FrequencyPenalty int      `json:"frequency_penalty"`
	Grammar          string   `json:"grammar"`
	ImageData        []any    `json:"image_data"`
	MinP             float64  `json:"min_p"`
	Mirostat         int      `json:"mirostat"`
	MirostatEta      float64  `json:"mirostat_eta"`
	MirostatTau      int      `json:"mirostat_tau"`
	NPredict         int      `json:"n_predict"`
	NProbs           int      `json:"n_probs"`
	PresencePenalty  int      `json:"presence_penalty"`
	Prompt           string   `json:"prompt"`
	RepeatLastN      int      `json:"repeat_last_n"`
	RepeatPenalty    float64  `json:"repeat_penalty"`
	SlotID           int      `json:"slot_id"`
	Stop             []string `json:"stop"`
	Stream           bool     `json:"stream"`
	Temperature      float64  `json:"temperature"`
	TfsZ             int      `json:"tfs_z"`
	TopK             int      `json:"top_k"`
	TopP             float64  `json:"top_p"`
	TypicalP         int      `json:"typical_p"`
}
*/

type LlamaResponse struct {
	Content            string `json:"content"`
	GenerationSettings struct {
		FrequencyPenalty float64       `json:"frequency_penalty"`
		Grammar          string        `json:"grammar"`
		IgnoreEos        bool          `json:"ignore_eos"`
		LogitBias        []interface{} `json:"logit_bias"`
		MinP             float64       `json:"min_p"`
		Mirostat         int           `json:"mirostat"`
		MirostatEta      float64       `json:"mirostat_eta"`
		MirostatTau      float64       `json:"mirostat_tau"`
		Model            string        `json:"model"`
		NCtx             int           `json:"n_ctx"`
		NKeep            int           `json:"n_keep"`
		NPredict         int           `json:"n_predict"`
		NProbs           int           `json:"n_probs"`
		PenalizeNl       bool          `json:"penalize_nl"`
		PresencePenalty  float64       `json:"presence_penalty"`
		RepeatLastN      int           `json:"repeat_last_n"`
		RepeatPenalty    float64       `json:"repeat_penalty"`
		Seed             int64         `json:"seed"`
		Stop             []interface{} `json:"stop"`
		Stream           bool          `json:"stream"`
		Temp             float64       `json:"temp"`
		TfsZ             float64       `json:"tfs_z"`
		TopK             int           `json:"top_k"`
		TopP             float64       `json:"top_p"`
		TypicalP         float64       `json:"typical_p"`
	} `json:"generation_settings"`
	Model        string `json:"model"`
	Prompt       string `json:"prompt"`
	SlotID       int    `json:"slot_id"`
	Stop         bool   `json:"stop"`
	StoppedEos   bool   `json:"stopped_eos"`
	StoppedLimit bool   `json:"stopped_limit"`
	StoppedWord  bool   `json:"stopped_word"`
	StoppingWord string `json:"stopping_word"`
	Timings      struct {
		PredictedMs         float64 `json:"predicted_ms"`
		PredictedN          int     `json:"predicted_n"`
		PredictedPerSecond  float64 `json:"predicted_per_second"`
		PredictedPerTokenMs float64 `json:"predicted_per_token_ms"`
		PromptMs            float64 `json:"prompt_ms"`
		PromptN             int     `json:"prompt_n"`
		PromptPerSecond     float64 `json:"prompt_per_second"`
		PromptPerTokenMs    float64 `json:"prompt_per_token_ms"`
	} `json:"timings"`
	TokensCached    int  `json:"tokens_cached"`
	TokensEvaluated int  `json:"tokens_evaluated"`
	TokensPredicted int  `json:"tokens_predicted"`
	Truncated       bool `json:"truncated"`
}

func NewLlamaChat(systemMsg string, backend string) *LlamaChat {
	cs := &LlamaChat{
		Prompt: systemMsg,
	}
	return cs
}

/*
	func (cs *LlamaChat) PreparePrompt() string {
		var sb strings.Builder
		for _, msg := range cs.messages {
			switch msg.Role {
			case openai.ChatMessageRoleUser:
				sb.WriteString("### User:\n")
			case openai.ChatMessageRoleAssistant:
				sb.WriteString("### Assistant:\n")
			case openai.ChatMessageRoleSystem:
				sb.WriteString("### System:\n")
			default:
				panic("Unknown role")
			}
			sb.WriteString(msg.Content)
			sb.WriteString("\n")
		}
		sb.WriteString("### Assistant:\n")
		return sb.String()
	}

	func (cs *LlamaChat) PrepareLLamaPrompt() string {
		var sb strings.Builder
		for _, msg := range cs.messages {
			switch msg.Role {
			case openai.ChatMessageRoleUser:
				sb.WriteString("[INST] <<SYS>> <</SYS>>\n\n" + msg.Content + " [/INST]")
			case openai.ChatMessageRoleAssistant:
				sb.WriteString("[INST] <<SYS>> <</SYS>>\n\n" + msg.Content + " [INST]")
			case openai.ChatMessageRoleSystem:
				sb.WriteString("[INST]<<SYS>>\n" + msg.Content + "\n<</SYS>>")
			default:
				panic("Unknown role")
			}
		}
		sb.WriteString("[INST] <<SYS>> <</SYS>>\n\n")
		return sb.String()
	}

	func (cs *LlamaChat) PreparePromptChatMLV1() string {
		var sb strings.Builder
		for _, msg := range cs.messages {
			sb.WriteString("<|im_start|>")
			sb.WriteString(msg.Role)
			sb.WriteString("\n")
			sb.WriteString(msg.Content)
			sb.WriteString("<|im_end|>\n")
		}
		sb.WriteString("<|im_start|>")
		sb.WriteString(openai.ChatMessageRoleAssistant)
		sb.WriteString("\n")
		return sb.String()
	}
*/

func (cs *LlamaChat) PreparePhi3Prompt(ch *ChatHistory) string {
	var sb strings.Builder
	//sb.WriteString("<|endoftext|>\n")
	sb.WriteString("<|system|>\n")
	sb.WriteString(cs.Prompt)
	sb.WriteString("<|end|>\n")
	for _, msg := range ch.Messages {
		switch msg.Role {
		case ChatHistoryRoleUser:
			sb.WriteString("<|user|>\n")
			sb.WriteString(msg.Content)
			sb.WriteString("<|end|>\n")
		case ChatHistoryRoleAssistant:
			sb.WriteString("<|assistant|>\n")
			sb.WriteString(msg.Content)
			sb.WriteString("<|end|>\n")
		default:
			panic("Unknown role")
		}
	}
	sb.WriteString("<|assistant|>\n")
	return sb.String()
}

func (cs *LlamaChat) PrepareLlama3Prompt(ch *ChatHistory) string {
	var sb strings.Builder
	sb.WriteString("<|begin_of_text|><|start_header_id|>system<|end_header_id|>\n\n")
	sb.WriteString(cs.Prompt)
	sb.WriteString("<|eot_id|>")

	for _, msg := range ch.Messages {
		switch msg.Role {
		case ChatHistoryRoleUser:
			sb.WriteString("<|start_header_id|>user<|end_header_id|>\n\n")
			sb.WriteString(msg.Content)
			sb.WriteString("<|eot_id|>")
		case ChatHistoryRoleAssistant:
			sb.WriteString("<|start_header_id|>assistant<|end_header_id|>\n\n")
			sb.WriteString(msg.Content)
			sb.WriteString("<|eot_id|>")
		default:
			panic("Unknown role")
		}
	}
	sb.WriteString("<|start_header_id|>assistant<|end_header_id|>\n\n")
	return sb.String()
}

func (cs *LlamaChat) GetResponse(ch *ChatHistory) (string, int, int) {

	req := LlamaRequest{
		//Prompt: cs.PreparePhi3Prompt(ch),
		Prompt: cs.PrepareLlama3Prompt(ch),
		// Prompt: cs.PreparePrompt(),
		// Prompt:      cs.PrepareLLamaPrompt(),
		NPredict:    1024,
		Stream:      false,
		CachePrompt: true,
		Stop: []string{
			"<|im_end|>",     // ChatMl
			"<|end|>",        // Phi3
			"### Assistant:", // Orca Hashes
			"### User:",      // Orca Hashes
			"<|eot_id|>",     // Llama3
			"[INST]"},
	}

	reqAsJson, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:8080/completion", "application/json", bytes.NewBuffer(reqAsJson))
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic(err)
	}
	defer resp.Body.Close()

	var response LlamaResponse
	err = json.NewDecoder(resp.Body).Decode(&response)

	return response.Content, response.TokensEvaluated, response.TokensPredicted

}

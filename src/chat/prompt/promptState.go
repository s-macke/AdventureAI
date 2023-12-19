package prompt

type Command struct {
	Situation string `json:"situation"`
	Narrator  string `json:"narrator"`
	Thought   string `json:"thought"`
	Command   string `json:"command"`
}

type State interface {
	ParseResponse(content string) Command
	GetSystemPrompt() string
}

package promptPattern

import (
	"github.com/s-macke/AdventureAI/src/chat/backend"
	"github.com/s-macke/AdventureAI/src/chat/storyHistory"
	"regexp"
	"strings"
)

type ReAct struct {
	re         *regexp.Regexp
	re2        *regexp.Regexp
	chatClient backend.ChatBackend
	systemMsg  string
}

type Command struct {
	Summary   string `json:"summary"`
	Situation string `json:"situation"`
	Narrator  string `json:"narrator"`
	Thought   string `json:"thought"`
	Command   string `json:"command"`
}

func NewPromptReAct(backendAsString string) *ReAct {
	// Your name is not Brian Hadley. You have accidentally killed Brian Hadley in the house.
	// Your first task is to look under your bed.`
	// You are a murderer.
	const systemMsg string = `You act as a player of an interactive text adventure. The goal is to win the game. 
The user provides the text of the text adventure. He is not a human and just prints the output of the game.

The format of your output must be:
SITUATION: {A short description of the current situation you are in.}
THOUGHT: {A curious, adventurous thought.}
COMMAND: {The single two word command you want to execute.}
`
	return &ReAct{
		chatClient: backend.NewChatBackend(systemMsg, backendAsString),
		re:         regexp.MustCompile(`\r?\n`),
		re2:        regexp.MustCompile(`SITUATION:(.*)THOUGHT:(.*)COMMAND:(.*)`),
		systemMsg:  systemMsg,
	}
}

func (c *ReAct) GetPrompt() string {
	return c.systemMsg
}

func (c *ReAct) GetNextCommand(story *storyHistory.StoryHistory) string {
	content, _, _ := c.chatClient.GetResponse(ToChatHistory(story))
	CheckAndShowContent(&content)

	temp := c.re.ReplaceAllString(content, " ")
	matches := c.re2.FindStringSubmatch(temp)
	cmd := Command{}
	if len(matches) < 4 { // fix for llama3 70B
		content = strings.ToLower(content)
		if content == "restart" {
			cmd.Command = "restart"
		} else if content == "quit" {
			cmd.Command = "quit"
		} else if content == "yes" {
			cmd.Command = "yes"
		} else if content == "no" {
			cmd.Command = "yes"
		} else {
			panic("Invalid response from backend")
		}
	} else {
		cmd.Situation = strings.TrimSpace(matches[1])
		//cmd.Narrator = strings.TrimSpace(matches[2])
		cmd.Thought = strings.TrimSpace(matches[2])
		cmd.Command = strings.TrimSpace(matches[3])
		if cmd.Command[0] == '"' && cmd.Command[len(cmd.Command)-1] == '"' {
			cmd.Command = cmd.Command[1 : len(cmd.Command)-1]
		}
		cmd.Command = strings.ReplaceAll(cmd.Command, ".", "")
	}
	story.AppendMessage(storyHistory.StoryMessage{
		Role:             "assistant",
		Content:          cmd.Command,
		CompletionTokens: 0,
		PromptTokens:     0,
		Meta:             content,
		Score:            -1,
	})

	return cmd.Command
}

func ToChatHistory(story *storyHistory.StoryHistory) *backend.ChatHistory {
	ch := backend.ChatHistory{
		Messages: []backend.ChatMessage{},
	}

	for _, msg := range story.Messages {
		if msg.Meta != "" {
			ch.Messages = append(ch.Messages, backend.ChatMessage{
				Role:    msg.Role,
				Content: msg.Meta,
			})
		} else {
			ch.Messages = append(ch.Messages, backend.ChatMessage{
				Role:    msg.Role,
				Content: msg.Content,
			})
		}
	}
	return &ch
}

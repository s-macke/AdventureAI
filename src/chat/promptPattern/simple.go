package promptPattern

import "C"
import (
	"github.com/s-macke/AdventureAI/src/chat/backend"
	"github.com/s-macke/AdventureAI/src/chat/storyHistory"
	"regexp"
	"strings"
)

type Simple struct {
	re         *regexp.Regexp
	chatClient backend.ChatBackend
	systemMsg  string
}

func NewPromptSimple(backendAsString string) *Simple {
	const systemMsg string = `You act as a player of an interactive text adventure. The goal is to win the game. 
The user provides the text of the text adventure. He is not a human and just prints the output of the game.

Your output should be a simple command, typically one or two words. 
`
	// The format of your output must be a single short command you want to execute.
	return &Simple{
		chatClient: backend.NewChatBackend(systemMsg, backendAsString),
		re:         regexp.MustCompile(`\r?\n`),
		systemMsg:  systemMsg,
	}
}

func (c *Simple) GetPrompt() string {
	return c.systemMsg
}

func (c *Simple) GetNextCommand(story *storyHistory.StoryHistory) string {
	content, _, _ := c.chatClient.GetResponse(ToChatHistory(story))
	CheckAndShowContent(&content)

	content = c.re.ReplaceAllString(content, " ")
	cmd := strings.TrimSpace(content)
	if cmd[0] == '"' && cmd[len(cmd)-1] == '"' {
		cmd = cmd[1 : len(cmd)-1]
	}
	cmd = strings.ReplaceAll(cmd, ".", "")

	// don't allow to quit
	if strings.ToLower(cmd) == "quit" {
		cmd = "wait"
		if strings.Contains(story.GetLastMessage().Content, "Would you like to RESTART") {
			cmd = "restart"
		}
	}
	story.AppendMessage(storyHistory.StoryMessage{
		Role:             "assistant",
		Content:          cmd,
		CompletionTokens: 0,
		PromptTokens:     0,
		Meta:             "",
		Score:            -1,
	})

	return cmd
}

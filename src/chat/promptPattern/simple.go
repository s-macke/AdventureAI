package promptPattern

import "C"
import (
	"fmt"
	"github.com/s-macke/AdventureAI/src/chat/backend"
	"github.com/s-macke/AdventureAI/src/chat/storyHistory"
	"regexp"
	"strings"
)

type Simple struct {
	re         *regexp.Regexp
	chatClient backend.ChatBackend
}

func NewPromptSimple(backendAsString string) *Simple {
	const systemMsg string = `You act as a player of an interactive text adventure. The goal is to win the game. 
The user provides the text of the text adventure. He is not a human and just prints the output of the game.

The format of your output must be a single two word command you want to execute. 
`

	return &Simple{
		chatClient: backend.NewChatBackend(systemMsg, backendAsString),
		re:         regexp.MustCompile(`\r?\n`),
	}
}

func (c *Simple) GetNextCommand(story *storyHistory.StoryHistory) (string, string) {
	content, _, _ := c.chatClient.GetResponse(story.GetLastMessage().Content)
	if content == "" {
		panic("empty content")
	}
	fmt.Printf(InfoColor, content)

	content = c.re.ReplaceAllString(content, " ")
	cmd := strings.TrimSpace(content)
	if cmd[0] == '"' && cmd[len(cmd)-1] == '"' {
		cmd = cmd[1 : len(cmd)-1]
	}
	cmd = strings.ReplaceAll(cmd, ".", "")
	return cmd, ""
}

package prompt

import "C"
import (
	"regexp"
	"strings"
)

type Simple struct {
	re *regexp.Regexp
}

func NewPromptSimple() *Simple {
	return &Simple{
		re: regexp.MustCompile(`\r?\n`),
	}
}

func (c *Simple) GetSystemPrompt() string {
	const systemMsg string = `You act as a player of an interactive text adventure. The goal is to win the game. 
The user provides the text of the text adventure. He is not a human and just prints the output of the game.

The format of your output must be a single two word command you want to execute. 
`
	return systemMsg
}

func (c *Simple) ParseResponse(content string) Command {
	cmd := Command{}

	content = c.re.ReplaceAllString(content, " ")
	cmd.Command = strings.TrimSpace(content)
	if cmd.Command[0] == '"' && cmd.Command[len(cmd.Command)-1] == '"' {
		cmd.Command = cmd.Command[1 : len(cmd.Command)-1]
	}
	cmd.Command = strings.ReplaceAll(cmd.Command, ".", "")
	return cmd
}

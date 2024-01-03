package prompt

import (
	"regexp"
	"strings"
)

type ReAct struct {
}

func NewPromptReAct() *ReAct {
	return &ReAct{}
}

func (c *ReAct) GetSystemPrompt() string {
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
	return systemMsg
}

func (c *ReAct) ParseResponse(content string) Command {
	cmd := Command{}
	re := regexp.MustCompile(`\r?\n`)
	content = re.ReplaceAllString(content, " ")
	re = regexp.MustCompile(`SITUATION:(.*)THOUGHT:(.*)COMMAND:(.*)`)
	matches := re.FindStringSubmatch(content)

	cmd.Situation = strings.TrimSpace(matches[1])
	//cmd.Narrator = strings.TrimSpace(matches[2])
	cmd.Thought = strings.TrimSpace(matches[2])
	cmd.Command = strings.TrimSpace(matches[3])
	if cmd.Command[0] == '"' && cmd.Command[len(cmd.Command)-1] == '"' {
		cmd.Command = cmd.Command[1 : len(cmd.Command)-1]
	}
	cmd.Command = strings.ReplaceAll(cmd.Command, ".", "")
	return cmd
}

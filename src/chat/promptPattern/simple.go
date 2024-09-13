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

func NewPromptSimple(stateAsString string, backendAsString string) *Simple {
	systemMsg := ""
	switch stateAsString {
	case "simple":
		systemMsg = `You act as a player of an interactive text adventure game. The goal is to win the game.
The user provides the text of the text adventure, as well as the initial scenario.
The user is not a human and will only prints the output of the game.

Your task is to respond with simple commands, typically one or two words.
		`
	case "simple_with_examples":
		systemMsg = `You act as a player of an interactive text adventure game. The goal is to win the game. 
The user provides the text of the text adventure, as well as the initial scenario. 
The user is not a human and will only prints the output of the game.

Your task is to respond with simple commands, typically one or two words.

These commands can include:

    Movement: "go north", "move south", "enter cave"
    Interaction with objects: "take sword", "open door", "read book"
    Examination: "look around", "inspect chest", "examine statue"
    Use of items: "use key", "drink potion", "light torch"
    Communication: "talk to guard", "ask villager", "shout"
    Inventory: "inventory", "check bag", "show items"

Consider the following guidelines:

    Context: Respond to the provided text by making decisions based on the scenario described.
    Clarity: Ensure your commands are clear and directly related to the current situation in the game.
    Feedback: Adjust your commands based on the results and feedback given by the user.
`
	default:
		panic("Unknown prompt")
	}

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
	if len(cmd) > 1 {
		if cmd[0] == '"' && cmd[len(cmd)-1] == '"' {
			cmd = cmd[1 : len(cmd)-1]
		}
	}
	cmd = strings.ReplaceAll(cmd, ".", "")

	// don't allow to quit
	if strings.ToLower(cmd) == "quit" {
		cmd = "wait"
		if strings.Contains(story.GetLastMessage().Content, "Would you like to RESTART") {
			cmd = "restart"
		}
	}
	if strings.ToLower(cmd) == "help" { // prevent showing up the help menu
		cmd = "hint"
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

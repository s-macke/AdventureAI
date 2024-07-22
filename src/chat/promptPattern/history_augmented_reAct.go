package promptPattern

import (
	"github.com/s-macke/AdventureAI/src/chat/backend"
	"github.com/s-macke/AdventureAI/src/chat/storyHistory"
	"regexp"
	"strings"
)

type HistoryAugmentedReact struct {
	re         *regexp.Regexp
	re2        *regexp.Regexp
	chatClient backend.ChatBackend
	systemMsg  string
}

func NewPromptHistoryAugmentedReact(backendAsString string) *HistoryAugmentedReact {
	const systemMsg string = `You act as a player of an interactive text adventure. 
The user provides the current state of the game with the whole history of commands and the outcome. 
You task is to continue the game and to win the game.

The format of your output must be:
SUMMARY: {A summary of the whole story so far.}
SITUATION: {A short description of the current you are in.}
THOUGHT: {A curious, adventurous thought.}
COMMAND: {The single two word command you want to execute.}
`
	return &HistoryAugmentedReact{
		re:         regexp.MustCompile(`\r?\n`),
		re2:        regexp.MustCompile(`SUMMARY:(.*)SITUATION:(.*)THOUGHT:(.*)COMMAND:(.*)`),
		chatClient: backend.NewChatBackend(systemMsg, backendAsString),
		systemMsg:  systemMsg,
	}
}

func (c *HistoryAugmentedReact) GetPrompt() string {
	return c.systemMsg
}

func (c *HistoryAugmentedReact) GetNextCommand(story *storyHistory.StoryHistory) string {
	ch := backend.ChatHistory{
		Messages: []backend.ChatMessage{{
			Role:    "user",
			Content: story.GetStory(),
		}},
	}
	content, _, _ := c.chatClient.GetResponse(&ch)
	CheckAndShowContent(&content)
	temp := c.re.ReplaceAllString(content, " ")
	matches := c.re2.FindStringSubmatch(temp)

	cmd := Command{}
	cmd.Summary = strings.TrimSpace(matches[1])
	cmd.Situation = strings.TrimSpace(matches[2])
	//cmd.Narrator = strings.TrimSpace(matches[2])
	cmd.Thought = strings.TrimSpace(matches[3])
	cmd.Command = strings.TrimSpace(matches[4])
	if cmd.Command[0] == '"' && cmd.Command[len(cmd.Command)-1] == '"' {
		cmd.Command = cmd.Command[1 : len(cmd.Command)-1]
	}
	cmd.Command = strings.ReplaceAll(cmd.Command, ".", "")

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

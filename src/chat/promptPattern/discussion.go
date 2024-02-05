package promptPattern

import (
	"fmt"
	"github.com/s-macke/AdventureAI/src/chat/backend"
	"github.com/s-macke/AdventureAI/src/chat/storyHistory"
	"regexp"
	"strings"
)

type Discussion struct {
	re              *regexp.Regexp
	backendAsString string
}

func NewPromptDiscussion(backendAsString string) *Discussion {
	return &Discussion{
		backendAsString: backendAsString,
		re:              regexp.MustCompile(`\[\[(.*)]]`),
	}
}

func (c *Discussion) GetNextCommand(story *storyHistory.StoryHistory) (string, string) {
	const systemMsg string = `The user provides you with a text adventure story up to a given point.

You describe the discussion about how to continue the adventure. 
The discussion is between 3 characters.

1. Dave: An expert on playing and winning text adventures.
2. Silvia: A newbie to text adventures. But very curious.
3. Gretel: A smart AI. She is a bit shy and does not talk much.

At the end of the discussion, all agree on the next two word command to continue the game.
The format of command must be 
[[{Two word command}]]
`
	chatClient := backend.NewChatBackend(systemMsg, c.backendAsString)
	content, _, _ := chatClient.GetResponse(story.GetStory())
	if content == "" {
		panic("empty content")
	}
	content = strings.ReplaceAll(content, "\r\n", "\n")
	fmt.Printf(InfoColor, content)

	matches := c.re.FindStringSubmatch(content)
	command := matches[len(matches)-1]

	if command == "" {
		panic("empty command")
	}
	fmt.Println("command: ", command)
	return command, content
}

package promptPattern

import (
	"fmt"
	"github.com/s-macke/AdventureAI/src/chat/backend"
	"github.com/s-macke/AdventureAI/src/chat/storyHistory"
	"regexp"
)

type Discussion struct {
	re         *regexp.Regexp
	chatClient backend.ChatBackend
}

func NewPromptDiscussion(backendAsString string) *Discussion {
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
	return &Discussion{
		re:         regexp.MustCompile(`\[\[(.*)]]`),
		chatClient: backend.NewChatBackend(systemMsg, backendAsString),
	}
}

func (c *Discussion) GetNextCommand(story *storyHistory.StoryHistory) string {
	ch := backend.ChatHistory{
		Messages: []backend.ChatMessage{{
			Role:    "user",
			Content: story.GetStory(),
		}},
	}

	content, _, _ := c.chatClient.GetResponse(&ch)
	CheckAndShowContent(&content)

	matches := c.re.FindStringSubmatch(content)
	command := matches[len(matches)-1]

	if command == "" {
		panic("empty command")
	}
	story.AppendMessage(storyHistory.StoryMessage{
		Role:             "assistant",
		Content:          command,
		CompletionTokens: 0,
		PromptTokens:     0,
		Meta:             content,
	})

	fmt.Println("command: ", command)
	return command
}

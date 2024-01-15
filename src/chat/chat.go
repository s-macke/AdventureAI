package chat

import (
	"fmt"
	"github.com/s-macke/AdventureAI/src/chat/backend"
	"github.com/s-macke/AdventureAI/src/chat/prompt"
	"github.com/s-macke/AdventureAI/src/zmachine"
)

const (
	InfoColor = "\033[1;34m%s\033[0m"
)

type ChatState struct {
	prompt           prompt.State
	chatClient       backend.ChatBackend
	zm               *zmachine.ZMachine
	story            *StoryState
	output           string
	currentStoryStep int
}

func NewChatState(zm *zmachine.ZMachine, chatPrompt string, backendAsString string) *ChatState {
	cs := &ChatState{
		zm:               zm,
		story:            &StoryState{},
		output:           "",
		currentStoryStep: 0,
	}
	fmt.Println("Use prompt: ", chatPrompt)
	switch chatPrompt {
	case "simple":
		cs.prompt = prompt.NewPromptSimple()
	case "react":
		cs.prompt = prompt.NewPromptReAct()
	default:
		panic("Unknown prompt")
	}
	cs.story.Prompt = cs.prompt.GetSystemPrompt()
	fmt.Println("Use backend: ", backendAsString)
	cs.chatClient = backend.NewChatBackend(cs.story.Prompt, backendAsString)
	cs.zm.Input = cs.chatInput
	//LoadStoryFromFile(&cs.story, cs.zm.Name)
	return cs
}

func (cs *ChatState) chatInput() string {
	cs.story.Messages = append(cs.story.Messages, StoryMessage{
		Role:             "user",
		Content:          cs.output,
		CompletionTokens: 0,
		PromptTokens:     0,
		IsResponse:       false,
		Command:          prompt.Command{},
	})

	fmt.Print(cs.output)

	// just return to the previous state if we have commands left
	if (cs.currentStoryStep*2 + 1) < len(cs.story.Messages) {
		step := cs.story.Messages[cs.currentStoryStep*2+1]
		if step.IsResponse {
			//fmt.Printf(InfoColor, "REASONING: "+step.Messages[len(cs.messages)-1].Content)
			fmt.Printf(InfoColor, "COMMAND: "+step.Command.Command+"\n")
			cs.currentStoryStep++
			return step.Command.Command
		} else {
			//panic("not a response")
		}
	}

	content, promptTokens, completionTokens := cs.chatClient.GetResponse(cs.output)
	cs.zm.Output.Reset()
	cs.output = ""

	fmt.Printf(InfoColor, content)
	fmt.Println()

	if content == "" {
		panic("empty content")
	}
	cmd := cs.prompt.ParseResponse(content)
	if cmd.Command == "" {
		panic("empty command")
	}

	cs.story.Messages = append(cs.story.Messages, StoryMessage{
		Role:             "assistant",
		Content:          cs.output,
		CompletionTokens: completionTokens,
		PromptTokens:     promptTokens,
		IsResponse:       true,
		Command:          cmd,
	})

	StoreStoryToFile(cs.story, cs.zm.Name)
	cs.currentStoryStep++

	if cs.currentStoryStep%5 == 0 {
		fmt.Println("Press ENTER to continue...")
		_, _ = fmt.Scanln()
	}

	return cmd.Command
}

func (cs *ChatState) ChatLoop() {
	for !cs.zm.Done {
		cs.zm.InterpretInstruction()
		if cs.zm.Output.Len() > 0 {
			if cs.zm.WindowId == 0 {
				cs.output += cs.zm.Output.String()
			}
			cs.zm.Output.Reset()
		}
	}
}

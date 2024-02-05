package chat

import (
	"fmt"
	"github.com/s-macke/AdventureAI/src/chat/promptPattern"
	"github.com/s-macke/AdventureAI/src/chat/storyHistory"
	"github.com/s-macke/AdventureAI/src/zmachine"
)

type ChatState struct {
	prompt           promptPattern.State
	zm               *zmachine.ZMachine
	story            *storyHistory.StoryHistory
	output           string
	currentStoryStep int
}

var commands = []string{
	"answer phone",
	"stand",
	"s",
	"remove watch",
	"remove clothes",
	"drop all",
	"enter shower",
	"take watch",
	"wear watch",
	"n",
	"get all from table",
	"open dresser",
	"get clothes",
	"wear clothes",
	"e",
	"open front door",
	"s",
	"open car with keys",
	"enter car",
	"no",
	"yes",
	"open wallet",
	"take ID",
	"insert card in slot",
	"enter cubicle",
	"read note",
	"take form and pen",
	"sign form",
	"out",
	"west",
	"restart",
	//"look under bed",
}

func NewChatState(zm *zmachine.ZMachine, chatPromptPattern string, backendAsString string) *ChatState {
	cs := &ChatState{
		zm: zm,
		story: &storyHistory.StoryHistory{
			PromptPattern: chatPromptPattern,
		},
		output:           "",
		currentStoryStep: 0,
	}

	fmt.Println("Use prompt: ", chatPromptPattern)
	cs.prompt = promptPattern.NewPrompt(chatPromptPattern, backendAsString)

	fmt.Println("Use backend: ", backendAsString)
	cs.zm.Input = cs.chatInput
	//cs.story.LoadFromFile(cs.zm.Name)
	return cs
}

func (cs *ChatState) chatInput() string {
	//cs.output does contain the output of the game of the current step
	fmt.Print(cs.output)

	cs.story.AppendMessage(storyHistory.StoryMessage{
		Role:             "user",
		Content:          cs.output,
		CompletionTokens: 0,
		PromptTokens:     0,
	})
	/*
		if cs.currentStoryStep < len(commands) {
			cs.currentStoryStep++
			cs.zm.Output.Reset()
			cs.output = ""

			cs.story.AppendMessage(storyHistory.StoryMessage{
				Role:             "assistant",
				Content:          commands[cs.currentStoryStep-1],
				CompletionTokens: 0,
				PromptTokens:     0,
				Meta:             "",
			})
			fmt.Printf("\033[1;34m%s\033[0m\n", commands[cs.currentStoryStep-1])
			return commands[cs.currentStoryStep-1]
		}
	*/
	// if we have a loaded history, return the next command
	if cmd, ok := cs.IsCommandStored(); ok {
		return cmd
	}

	cmd, meta := cs.prompt.GetNextCommand(cs.story)
	if cmd == "" {
		panic("empty command")
	}

	cs.zm.Output.Reset()
	cs.output = ""

	cs.story.AppendMessage(storyHistory.StoryMessage{
		Role:             "assistant",
		Content:          cmd,
		CompletionTokens: 0,
		PromptTokens:     0,
		Meta:             meta,
	})

	cs.story.StoreToFile(cs.zm.Name)
	cs.currentStoryStep++

	if cs.currentStoryStep%5 == 0 {
		fmt.Println("Press ENTER to continue...")
		_, _ = fmt.Scanln()
	}

	return cmd
}

func (cs *ChatState) IsCommandStored() (string, bool) {
	// just return to the previous state if we have commands left
	if (cs.currentStoryStep*2 + 1) < len(cs.story.Messages) {
		step := cs.story.Messages[cs.currentStoryStep*2+1]
		if step.Role == "assistant" {
			fmt.Printf("%s\n", "COMMAND: "+step.Content)
			cs.currentStoryStep++
			return step.Content, true
		} else {
			//panic("not a response")
		}
	}
	return "", false
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

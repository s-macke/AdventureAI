package mainsrc

import (
	"bufio"
	"github.com/s-macke/AdventureAI/src/chat"
	"os"
)

func Input() string {
	/*
		input, ok := getWalkthrough(filename)
		if ok {
			return input
		}
	*/
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func Main() {
	config := parseConfig()
	zm := Init(config.filename)

	if config.mcp {
		// zm is not used in MCP mode, so we can ignore it. But the existence of the file is checked.
		StartServer(config.filename)
		return
	}

	if config.doChat {
		chatState := chat.NewChatState(zm, config.prompt, config.backend, config.oldStoryFilename)
		chatState.ChatLoop()
		return
	}

	zm.Input = Input
	for !zm.Done {
		zm.InterpretInstruction()
		if zm.Output.Len() > 0 {
			if zm.WindowId == 0 {
				_, _ = os.Stdout.WriteString(zm.Output.String())
				//fmt.Println("Score: ", zm.ReadGlobal(59)) // Score for Suveh Nux
			}
			zm.Output.Reset()
		}
	}
}

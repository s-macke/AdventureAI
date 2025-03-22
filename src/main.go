package mainsrc

import (
	"bufio"
	"github.com/s-macke/AdventureAI/src/chat"
	"github.com/s-macke/AdventureAI/src/zmachine"
	"os"
	"path/filepath"
	"strconv"
)

func Init(filename string) *zmachine.ZMachine {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Read %d bytes\n", len(buffer))

	var header zmachine.ZHeader
	header.Read(buffer)

	if header.Version != 3 && header.Version != 4 && header.Version != 5 && header.Version != 8 {
		panic("Only Version 3, 4, 5 or 8 files supported. But found version " + strconv.Itoa(int(header.Version)))
	}

	zm := zmachine.NewZMachine(filepath.Base(filename), buffer, header)
	return zm
}

var filename = ""

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

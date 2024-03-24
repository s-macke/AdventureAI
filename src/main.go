package mainsrc

import (
	"bufio"
	"flag"
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

	if header.Version != 3 && header.Version != 5 {
		panic("Only Version 3 and 5 files supported. But found version " + strconv.Itoa(int(header.Version)))
	}

	zm := zmachine.NewZMachine(filepath.Base(filename), buffer, header)
	return zm
}

/*
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
	"RESTART",
	"look under bed",
	"look at corpse",
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
	"yes",
	"no",
	"yes",
}

var commandIndex = 0
*/

func Input() string {
	/*
		if commandIndex < len(commands) {
			fmt.Println(commands[commandIndex])
			commandIndex++
			return commands[commandIndex-1]
		}
	*/
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func Main() {
	filename := flag.String("file", "905.z5", "Z-Machine file to run")
	doChat := flag.Bool("ai", false, "Chat with AI")
	prompt := flag.String("prompt", "react", "Chat with AI via prompt 'simple', or 'discuss', 'react' (reason and act) or 'history_react'")
	backend := flag.String("backend", "gpt4", "Select AI backend. Either 'gpt3', 'gpt4', 'orca2', 'mistral', 'gemini', 'claude'")
	flag.Parse()

	zm := Init(*filename)

	if *doChat {
		chatState := chat.NewChatState(zm, *prompt, *backend)
		chatState.ChatLoop()
		return
	}

	zm.Input = Input
	for !zm.Done {
		zm.InterpretInstruction()
		if zm.Output.Len() > 0 {
			if zm.WindowId == 0 {
				_, _ = os.Stdout.WriteString(zm.Output.String())
			}
			zm.Output.Reset()
		}
	}
}

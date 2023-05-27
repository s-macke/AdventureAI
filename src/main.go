package zmachine

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func Init(filename string) *ZMachine {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Read %d bytes\n", len(buffer))

	var header ZHeader
	header.read(buffer)

	if header.version != 3 && header.version != 5 {
		panic("Only Version 3 and 5 files supported. But found version " + strconv.Itoa(int(header.version)))
	}

	zm := NewZMachine(buffer, header)
	return zm
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
}
var commandIndex = 0

func Input() string {
	if commandIndex < len(commands) {
		fmt.Println(commands[commandIndex])
		commandIndex++
		return commands[commandIndex-1]
	}
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func Main() {
	filename := flag.String("file", "905.z5", "Z-Machine file to run")
	doChat := flag.Bool("ai", false, "Chat with AI")
	flag.Parse()

	zm := Init(*filename)

	if *doChat {
		chat(zm)
		return
	}

	zm.input = Input
	for !zm.done {
		zm.InterpretInstruction()
		if zm.output.Len() > 0 {
			if zm.windowId == 0 {
				_, _ = os.Stdout.WriteString(zm.output.String())
			}
			zm.output.Reset()
		}
	}

}

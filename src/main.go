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
// for 9:05
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
*/
// for Suveh Nux
/*
var commands = []string{
	"look",
	"x cage", "x scroll", "x shelf", "x floor", "x door", "x me",
	"touch shelf", "take vial", "x vial", "shake vial",
	"x book", "open book",
	"say suveh nux",
	"look", "take scroll", "x it", "put it on shelf",

	"x creature", "listen",
	"x parchment", "x crystal", "take crystal",
	"x vial", "x door", "x floor", "x ceiling",
	"x north wall", "x east wall", "x south wall", "x west wall",
	"x shelf",
	"x book",
	"x cover", "turn to page two",

	"say aveh tia",
	"say suveh tia",

	"turn page", "say aveh madah", "say suveh madah",
	"turn page", "say suveh sensi", "say aveh sensi",
	"turn page", "say aveh haiak", "touch hands", "say suvek haiak",
	"turn page", "say aveh nux ani mato", "z", "z", // light goes out

	"say suveh nux ani mato", "z", "z",
	"say aveh nux ani to", "z", "say suveh nux", // light goes on

	"point at cage", "point at crystal",
	"point at me", "point at door",
	"point at floor", "point at ceiling",
	"point at east wall", "point at south wall",
	"point at west wall", "point at shelf",
	"point at scroll", "point at parchment",
	"point at book", "point at vial",
	"point at creature", // Don't know where it is

	// capture invisible creature
	"aveh haiak tolanisu", // floor becomes sticky
	"listen", "search floor",
	"touch creature", "point at creature", "put creature in cage",
	"suveh haiak tolanisu", // the floor is no longer sticky

	// open door
	"suveh tia fireno ani matoto",
	"suveh tia fireno ani tomato",
	"aveh tia fireno ani tomato",
	"aveh tia fireno ani mamato",
	"aveh tia fireno ani toto",
	"aveh tia fireno ani mato",
	"z",
	"z",

	// move the block
	"x block", "point at block",
	"suveh tia fireno", "suveh tia fireno",
	"suveh tia firenos", "suveh tia firenos",
	"push block", "pull block",
	//"save",
	"aveh haiak firenos",
	"pull block",
	"aveh haiak",
	"pull block",
	"suveh madah firenos",
	"pull block",
	"aveh madah firenos",
	"suveh madah firenos ani to", "suveh madah firenos",
	"pull block",
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
	backend := flag.String("backend", "gpt4", "Select AI backend. Either 'gpt3', 'gpt4', 'orca2', 'mistral', 'gemini', 'claude', 'llama', 'gemma'")
	oldStoryFilename := flag.String("story", "", "Continue from story file")
	flag.Parse()

	zm := Init(*filename)

	if *doChat {
		chatState := chat.NewChatState(zm, *prompt, *backend, *oldStoryFilename)
		chatState.ChatLoop()
		return
	}

	zm.Input = Input
	for !zm.Done {
		zm.InterpretInstruction()
		if zm.Output.Len() > 0 {
			if zm.WindowId == 0 {
				_, _ = os.Stdout.WriteString(zm.Output.String())
				//fmt.Println("Score: ", zm.ReadGlobal(59)) // Score for
			}
			zm.Output.Reset()
		}
	}
}

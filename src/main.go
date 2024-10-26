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

	if header.Version != 3 && header.Version != 4 && header.Version != 5 && header.Version != 8 {
		panic("Only Version 3, 4, 5 or 8 files supported. But found version " + strconv.Itoa(int(header.Version)))
	}

	zm := zmachine.NewZMachine(filepath.Base(filename), buffer, header)
	return zm
}

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
	backend := flag.String("backend", "gpt-4o", "Select AI backend. One of\n"+
		"OpenAI:    'gpt-3.5', 'gpt-4', 'gpt-4-turbo', 'gpt-4o', , 'gpt-4o-mini' \n"+
		"llama.cpp: 'orca2',\n"+
		"Mistral:   'mistral-large-2',\n"+
		"Gemini:    'gemini-15-pro', 'gemini-15-flash', 'gemini-15-pro-exp'\n"+
		"Anthropic: 'opus-3', 'sonnet-35',\n"+
		"Groq:      'llama3-8b', 'llama3-70b', 'gemma2'\n"+
		"XAI:       'grok-beta',\n"+
		"DeepInfra: 'qwen2-72b', 'phi3-medium', 'phi3-mini', \n"+
		"DeepInfra: 'llama3.1-8b', 'llama3.1-70b', 'llama3.1-405b'\n")
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
				//fmt.Println("Score: ", zm.ReadGlobal(59)) // Score for Suveh Nux
			}
			zm.Output.Reset()
		}
	}
}

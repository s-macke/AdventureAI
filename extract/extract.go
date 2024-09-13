package main

import (
	"flag"
	"fmt"
	"os"
)

func PrintStory(state *StoryHistory) {
	// open input file
	fi, err := os.Create("output.md")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	_, _ = fmt.Fprintln(fi, "# "+state.Name)
	_, _ = fmt.Fprintln(fi, "")
	_, _ = fmt.Fprintln(fi, "* played with "+state.Model)
	_, _ = fmt.Fprintln(fi, "```")
	_, _ = fmt.Fprintln(fi, state.Prompt)
	_, _ = fmt.Fprintln(fi, "```")

	for _, message := range state.Messages {
		if message.Role == "user" {
			_, _ = fmt.Fprintln(fi, message.Content)
		}
		if message.Role == "assistant" {
			text := "> " + message.Content
			/*
				text := "> * " + message.Meta
				text = strings.Replace(text, "\n\n", "\n", -1)
				text = strings.Replace(text, "\n", "\n> * ", -1)
			*/
			_, _ = fmt.Fprintln(fi, text)
			_, _ = fmt.Fprintln(fi, "")
		}
		//_, _ = fmt.Fprintln(fi, step.Messages[len(step.Messages)-3].Content)
		//_, _ = fmt.Fprintln(fi, "> * **Situation:** "+step.Command.Situation)
		//_, _ = fmt.Fprintln(fi, "> * **Thought:** "+step.Command.Thought)
		//_, _ = fmt.Fprintln(fi, "> * **Command:** "+step.Command.Command)
		//_, _ = fmt.Fprintln(fi, "")
	}
}

func RenameModel(name string) string {
	switch name {
	case "sonnet-35":
		return "Claude 3.5 Sonnet"
	case "opus-3":
		return "Claude 3 Opus"
	case "gpt-4o-mini":
		return "GPT-4o mini"
	case "gpt-4o":
		return "GPT-4o"
	case "gpt-4":
		return "GPT-4"
	case "gpt-4-turbo":
		return "GPT-4 Turbo"
	case "llama3-70b":
		return "Llama 3 70B"
	case "llama3-8b":
		return "Llama 3 8B"
	case "llama3.1-8b":
		return "Llama 3.1 8B"
	case "llama3.1-70b":
		return "Llama 3.1 70B"
	case "llama3.1-405b":
		return "Llama 3.1 405B"
	case "gemma-2":
		return "Gemma 2"
	case "gemini-15-pro":
		return "Gemini 1.5 Pro"
	case "qwen2-72b":
		return "Qwen2 72B"
	case "phi3-medium":
		return "Phi 3 Medium"
	case "phi3-mini":
		return "Phi 3 Mini"
	}
	return name
}

func main() {
	filename := flag.String("file", "", "History File")
	progress := flag.Bool("progress", false, "prepare progress plot")
	flag.Parse()
	if *progress {
		PlotProgress("suvehnux")
		PlotProgress("905")
		PlotProgress("shade")
		PlotProgress("violet")
		PlotProgress("hhgg")
		//PlotProgress("ChildsPlay")
		os.Exit(0)
	}

	if *filename == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	var state StoryHistory
	if len(os.Args) < 2 {
		fmt.Println("Usage: extract <filename>")
		os.Exit(0)
	}
	LoadStoryFromFile(&state, *filename)
	//PrintPrices(&state)
	PrintStory(&state)
}

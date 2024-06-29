package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type StoryMessage struct {
	Role             string   `json:"role"`
	Content          string   `json:"content"`
	CompletionTokens int      `json:"completionTokens"`
	PromptTokens     int      `json:"promptTokens"`
	Meta             string   `json:"meta"`
	Score            *float64 `json:"score"`
}

type StoryHistory struct {
	PromptPattern string         `json:"promptPattern"`
	Model         string         `json:"model"`
	Prompt        string         `json:"prompt"`
	Date          string         `json:"date"`
	Name          string         `json:"name"`
	Messages      []StoryMessage `json:"messages"`
}

func LoadStoryFromFile(state *StoryHistory, name string) {
	filename := name
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &state)
	if err != nil {
		panic(err)
	}
}

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
			text := "> * " + message.Meta
			text = strings.Replace(text, "\n\n", "\n", -1)
			text = strings.Replace(text, "\n", "\n> * ", -1)

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

func PlotProgress() {
	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	fi, err := os.Create("progress/progress.dat")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	figp, err := os.Create("progress/plotlines.gp")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := figp.Close(); err != nil {
			panic(err)
		}
	}()

	fileidx := 0
	linetypeidx := 0
	lastmodel := ""
	_, _ = fmt.Fprintln(figp, "plot \\")
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), "suvehnux") || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		var state StoryHistory
		LoadStoryFromFile(&state, file.Name())
		_, _ = fmt.Fprintln(fi, "#", state.Model, state.PromptPattern)
		for i, message := range state.Messages {
			if message.Score != nil {
				_, _ = fmt.Fprintln(fi, i/2+1, float32(*message.Score)+float32(linetypeidx)*0.2)
			}
		}
		_, _ = fmt.Fprintln(fi, "")
		printModel := state.Model
		if lastmodel != state.Model {
			linetypeidx++
			lastmodel = state.Model
		} else {
			printModel = ""
		}
		_, _ = fmt.Fprintf(figp, "\"progress.dat\" every :::%d::%d w l ls %d title \"%s\", \\\n",
			fileidx, fileidx, linetypeidx,
			printModel)
		fileidx++
	}
}

func main() {
	filename := flag.String("file", "", "History File")
	progress := flag.Bool("progress", false, "prepare progress plot")
	flag.Parse()
	if *progress {
		PlotProgress()
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

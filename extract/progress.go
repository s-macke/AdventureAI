package main

import (
	"fmt"
	"os"
	"strings"
)

type StoryProgress struct {
	model string
	avg   [][]int
}

func (p *StoryProgress) GetAverage(idx int) float64 {
	sum := 0
	for _, v := range p.avg[idx] {
		sum += v
	}
	return float64(sum) / float64(len(p.avg[idx]))
}

func LoadProgress(data map[string]*StoryProgress, prefix string) {
	path := "../storydump/"
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if (!strings.HasPrefix(file.Name(), prefix+".z5_") && !strings.HasPrefix(file.Name(), prefix+".z8_")) || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		var state StoryHistory
		LoadStoryFromFile(&state, path+file.Name())
		if state.PromptPattern != "simple" {
			continue
		}
		if len(state.Messages) > 200 {
			state.Messages = state.Messages[:200]
		}

		fmt.Printf("Prepare %s %s %s %d\n", file.Name(), state.Model, state.PromptPattern, len(state.Messages))
		//name := state.Model + " " + state.PromptPattern

		detailedProgress := &StoryProgress{}
		detailedProgress.model = RenameModel(state.Model)
		if _, ok := data[detailedProgress.model]; ok {
			detailedProgress = data[detailedProgress.model]
		}

		pos := 0
		for _, message := range state.Messages {
			if message.Score != nil && *message.Score != -1 {
				if pos < len(detailedProgress.avg) {
					detailedProgress.avg[pos] = append(detailedProgress.avg[pos], int(*message.Score))
				} else {
					detailedProgress.avg = append(detailedProgress.avg, []int{int(*message.Score)})
				}
				pos++
			}
		}
		data[detailedProgress.model] = detailedProgress
	}
}

func PlotProgress(prefix string) {
	data := make(map[string]*StoryProgress)
	LoadProgress(data, prefix)
	sorted := SortMap(data)
	/*
		for k, v := range data {
			fmt.Println(k, v)
		}*/

	fi, err := os.Create("progress/" + prefix + "_progress.dat")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	figp, err := os.Create("progress/" + prefix + "_plotlines.gp")
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
	_, _ = fmt.Fprintln(figp, "plot \\")
	for i := range sorted {
		modelName := sorted[i]
		v := data[modelName]
		linetypeidx++
		_, _ = fmt.Fprintln(fi, "#", modelName)
		for i := range len(v.avg) {
			//_, _ = fmt.Fprintln(fi, i, float32(*message.Score)+float32(linetypeidx)*0.2)
			_, _ = fmt.Fprintln(fi, i, float32(v.GetAverage(i)))
		}
		_, _ = fmt.Fprintln(fi, "")
		//title := fmt.Sprintf("%s (%d)", modelName, len(v.avg[0]))
		title := fmt.Sprintf("%s", modelName)
		_, _ = fmt.Fprintf(figp, "\""+prefix+"_progress.dat\" every :::%d::%d w l ls %d title \"%s\", \\\n",
			fileidx, fileidx, linetypeidx, title)
		fileidx++
	}
}

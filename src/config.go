package mainsrc

import (
	"flag"
)

type config struct {
	filename         string
	doChat           bool
	prompt           string
	backend          string
	oldStoryFilename string
}

func parseConfig() config {
	filename := flag.String("file", "905.z5", "Z-Machine file to run")
	doChat := flag.Bool("ai", false, "Chat with AI")
	prompt := flag.String("prompt", "react", "Chat with AI via prompt 'simple', or 'discuss', 'react' (reason and act) or 'history_react'")
	backend := flag.String("backend", "gpt-4o", "Select AI backend. One of\n"+
		"OpenAI:    'gpt-3.5', 'gpt-4', 'gpt-4-turbo', 'gpt-4o', 'o1', 'o3-mini' \n"+
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

	return config{
		filename:         *filename,
		doChat:           *doChat,
		prompt:           *prompt,
		backend:          *backend,
		oldStoryFilename: *oldStoryFilename,
	}
}

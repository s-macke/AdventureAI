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
		"OpenAI:    'openai:gpt-3.5', 'openai:gpt-4', 'openai:gpt-4-turbo', 'openai:gpt-4o', 'openai:o3', 'openai:o4-mini' \n"+
		"llama.cpp: 'llama:orca2',\n"+
		"Mistral:   'mistral:mistral-large-2',\n"+
		"Gemini:    'gemini:gemini-2.5-pro', 'gemini:gemini-2.5-flash', 'gemini:gemini-2.5-flash-lite'\n"+
		"Anthropic: 'claude:opus-3', 'claude:sonnet-35', 'claude:sonnet-4', 'claude:opus-4',\n"+
		"Groq:      'groq:llama3-8b', 'groq:llama3-70b', 'groq:gemma2'\n"+
		"ollama:    'ollama:gemma3', 'ollama:qwen3-0.6b'\n"+
		"XAI:       'xai:grok-beta',\n"+
		"DeepInfra: 'deepinfra:qwen2-72b', 'deepinfra:phi3-medium', 'deepinfra:phi3-mini', \n"+
		"DeepInfra: 'deepinfra:llama3.1-8b', 'deepinfra:llama3.1-70b', 'deepinfra:llama3.1-405b'\n")
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

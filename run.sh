set -e

go build

./AdventureAI -ai -backend sonnet35 -prompt simple -file suvehnux.z5
#./AdventureAI -ai -backend opus3 -prompt simple -file suvehnux.z5
#./AdventureAI -ai -backend gemini15pro -prompt simple -file suvehnux.z5
#./AdventureAI -ai -backend gpt4 -prompt simple -file suvehnux.z5
#./AdventureAI -ai -backend gpt4o -prompt simple -file suvehnux.z5
#./AdventureAI -ai -backend gemma -prompt simple -file suvehnux.z5

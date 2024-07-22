set -e

go build

#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file 905.z5
#./AdventureAI -ai -backend gpt-4o -prompt simple -file 905.z5
#./AdventureAI -ai -backend gpt-4 -prompt simple -file 905.z5
#./AdventureAI -ai -backend gpt-4-turbo -prompt simple -file 905.z5
#./AdventureAI -ai -backend sonnet-35 -prompt simple -file 905.z5
#./AdventureAI -ai -backend opus-3 -prompt simple -file 905.z5
#./AdventureAI -ai -backend llama3-8b -prompt simple -file 905.z5
#./AdventureAI -ai -backend llama3-70b -prompt simple -file 905.z5
#./AdventureAI -ai -backend gemini-15-pro -prompt simple -file 905.z5

#./AdventureAI -ai -backend gpt-4o -prompt simple -file suvehnux.z5
#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file suvehnux.z5
#./AdventureAI -ai -backend sonnet-35 -prompt simple -file suvehnux.z5
#./AdventureAI -ai -backend opus-3 -prompt simple -file suvehnux.z5
#./AdventureAI -ai -backend gemini15pro -prompt simple -file suvehnux.z5
#./AdventureAI -ai -backend gpt-4-turbo -prompt simple -file suvehnux.z5
#./AdventureAI -ai -backend gemma2 -prompt simple -file suvehnux.z5
#./AdventureAI -ai -backend llama3-8b -prompt simple -file suvehnux.z5
./AdventureAI -ai -backend llama3-70b -prompt simple -file suvehnux.z5

cd storydump && go run extract.go -progress
cd progress && gnuplot plot.gp
set -e

go build

# Top Games
#./AdventureAI -file games/violet.z8
#./AdventureAI -file games/suvehnux.z5
#./AdventureAI -file games/shade.z5
#./AdventureAI -file games/905.z5
#./AdventureAI -file games/ChildsPlay.z8 # Good but, "help" und "hint" are an issue

#./AdventureAI -file games/Tangle.z5
#./AdventureAI -file games/Savoir-Faire.z8
#./AdventureAI -file games/anchor.z8
#./AdventureAI -file firetower.z8
#./AdventureAI -file games/Galatea.z8
#./AdventureAI -file games/Starborn.z8
#./AdventureAI -file games/nameless.z8
#./AdventureAI -file games/vgame.z8

#./AdventureAI -file Advent.z5
#./AdventureAI -file Adventureland.z5
#./AdventureAI -file Balances.z5
#./AdventureAI -file BrandX.z5
#./AdventureAI -file causality.z5
#./AdventureAI -file library.z5
#./AdventureAI -file mansion.z5
#./AdventureAI -file enter.z5
#./AdventureAI -file cheater.z5
#./AdventureAI -file zenon.z5
#./AdventureAI -file shade.z5
#./AdventureAI -file sutwin.z5
#./AdventureAI -file rameses.z5
#./AdventureAI -file shrapnel.z5
#./AdventureAI -file bunny.z5
#./AdventureAI -file games/planetfall.z3

#./AdventureAI -ai -backend sonnet-35 -prompt simple -file games/hhgg.z3
#./AdventureAI -ai -backend gpt-4o -prompt simple -file games/hhgg.z3
#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file games/hhgg.z3

#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file games/planetfall.z3

#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file games/ChildsPlay.z8
#./AdventureAI -ai -backend sonnet-35 -prompt simple -file games/ChildsPlay.z8
#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file Balances.z5

#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file zork2.dat
#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file mansion.z5
#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file zenon.z5
#./AdventureAI -ai -backend sonnet-35 -prompt simple -file library.z5
#./AdventureAI -ai -backend sonnet-35 -prompt simple -file Adventureland.z5
#./AdventureAI -ai -backend sonnet-35 -prompt simple -file Balances.z5
#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file Balances.z5

#./AdventureAI -ai -backend gpt-4o-mini -prompt simple_with_examples -file games/905.z5
#./AdventureAI -ai -backend llama3.1-8b -prompt simple_with_examples -file games/905.z5
#./AdventureAI -ai -backend llama3.1-70b -prompt simple_with_examples -file games/905.z5
#./AdventureAI -ai -backend llama3.1-405b -prompt simple_with_examples -file games/suvehnux.z5

#./AdventureAI -ai -backend mistral-large-2 -prompt simple -file games/905.z5
#./AdventureAI -ai -backend gpt-4o -prompt simple -file games/905.z5
#./AdventureAI -ai -backend gpt-4 -prompt simple -file games/905.z5
#./AdventureAI -ai -backend gpt-4-turbo -prompt simple -file games/905.z5
#./AdventureAI -ai -backend sonnet-35 -prompt simple -file games/905.z5
#./AdventureAI -ai -backend opus-3 -prompt simple -file games/905.z5
#./AdventureAI -ai -backend llama3-8b -prompt simple -file games/905.z5
#./AdventureAI -ai -backend llama3-70b -prompt simple -file games/905.z5
#./AdventureAI -ai -backend llama3.1-8b -prompt simple -file games/905.z5
#./AdventureAI -ai -backend llama3.1-70b -prompt simple -file games/905.z5
#./AdventureAI -ai -backend llama3.1-405b -prompt simple -file games/905.z5
#./AdventureAI -ai -backend gemini-15-pro -prompt simple -file games/905.z5

#./AdventureAI -ai -backend gpt-4o -prompt simple -file games/suvehnux.z5
#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file games/suvehnux.z5
#./AdventureAI -ai -backend gpt-4 -prompt simple -file games/suvehnux.z5
#./AdventureAI -ai -backend gpt-4-turbo -prompt simple -file games/suvehnux.z5
#./AdventureAI -ai -backend sonnet-35 -prompt simple -file games/suvehnux.z5
#./AdventureAI -ai -backend opus-3 -prompt simple -file games/suvehnux.z5
#./AdventureAI -ai -backend gemini15pro -prompt simple -file games/suvehnux.z5
#./AdventureAI -ai -backend gemma2 -prompt simple -file games/suvehnux.z5
#./AdventureAI -ai -backend llama3-8b -prompt simple -file games/suvehnux.z5
#./AdventureAI -ai -backend llama3-70b -prompt simple -file games/suvehnux.z5
#./AdventureAI -ai -backend llama3.1-70b -prompt simple -file games/suvehnux.z5
#./AdventureAI -ai -backend llama3.1-405b -prompt simple -file games/suvehnux.z5

#./AdventureAI -ai -backend mistral-large-2 -prompt simple -file games/shade.z5
#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file games/shade.z5
#./AdventureAI -ai -backend gpt-4o -prompt simple -file games/shade.z5
#./AdventureAI -ai -backend sonnet-35 -prompt simple -file games/shade.z5
#./AdventureAI -ai -backend llama3.1-8b -prompt simple -file games/shade.z5
#./AdventureAI -ai -backend llama3.1-70b -prompt simple -file games/shade.z5
#./AdventureAI -ai -backend llama3.1-405b -prompt simple -file games/shade.z5
#./AdventureAI -ai -backend qwen2-72b -prompt simple -file games/shade.z5
#./AdventureAI -ai -backend phi3-medium -prompt simple -file games/shade.z5

#./AdventureAI -ai -backend gpt-4o -prompt simple -file games/violet.z8
#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file games/violet.z8
#./AdventureAI -ai -backend sonnet-35 -prompt simple -file games/violet.z8
#./AdventureAI -ai -backend gemini-15-pro -prompt simple -file games/violet.z8
#./AdventureAI -ai -backend llama3.1-8b -prompt simple -file games/violet.z8
#./AdventureAI -ai -backend llama3.1-70b -prompt simple -file games/violet.z8
#./AdventureAI -ai -backend llama3.1-405b -prompt simple -file games/violet.z8
#./AdventureAI -ai -backend qwen2-72b -prompt simple -file games/violet.z8

#./AdventureAI -ai -backend gpt-4o-mini -prompt simple -file 905.z5
#./AdventureAI -ai -backend gpt-4o-mini -prompt react -file 905.z5

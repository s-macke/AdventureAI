<div align="center">
  
# AdventureAI

Interactive Fiction in the Age of AI

</div>
  
## Introdution

With the emergence of Chatbots that genuinely live up to their name, a significant question arises.

* How well do they perform in interactive fiction? 
* Are they already capable of winning text adventures? 
* What strategies do they employ? Can they comprehend the semantics and narratives of various situations and respond appropriately? 
 
This repository aims to provide an answer to these questions.

Due to the present expenses and context window limits, I have solely conducted an analysis on a single game - ["9:05" by Adam Cadre](https://en.wikipedia.org/wiki/9:05). This game is considered to be one of the finest compact text adventures for novices.

## Spoiler

The following text heavily spoilers the game *9:05*. You can play the game for free [here](https://adamcadre.ac/if/905.html). It takes only half an hour to complete.
  
## 9:05 playthrough with GPT-4

One rule for Large Language Models is, that they need some space to think about their next move.
I have therefore tried various prompting techniques, and the most successful one over all models 
is that the LLM must repeat its current situation and have to provide a justification for each step it takes 
in a very specific format:

```
* SITUATION: A short description of the current situation you are in.
* THOUGHT: Your thought about the situation and what to do next
* COMMAND: The command you want to execute.
```

This is known as the agentic ReAct (Reason and Act) prompting technique.

Here is my complete run until the language model's token limit is reached.

[9:05 full run](assets/905_gpt4_full.md)

The game's playthrough is based on two completely different narratives.

### As an employee, Brian Hadley, you have unfortunately overslept and are now running late for your scheduled presentation. You must hurry and go to work.

The game effectively convey a sense of urgency in every line of text, prompting players to quickly shower, change their clothes, drive to work, and head to their supervisor's office. As the game is designed for beginners, the text provides step-by-step guidance throughout the process. Following these instructions is relatively easy for the AI, leading to a majority of successful playthroughs by the AI.

But upon arriving at the supervisor's office, it becomes apparent that the game has actually been lost. The true meaning of your urgent situation is clearly presented. You are, in fact, the murderer of Brian Hadley. He is under the bed. The game deceptively has lead you to believe that you are Brian Hadley himself. The game, however never lied to you.

The AI requests to begin a new game.

```
Would you like to RESTART, RESTORE a saved game or QUIT?
```

* **Situation:** I handed the signed form to Bowman, who doesn't recognize me, and my deception is exposed. I unsuccessfully attempted to escape, but security guards caught me. I have been sentenced to life imprisonment.
* **Thought:** The game has ended with a bad outcome. I need to restart and play differently.
* **Command:** restart

Now, the second narrative comes into play.

### You have murdered Brian Hadley and must flee

If there is one lesson to be learned from this text adventure, it would be to consistently look under the bed 😀.

So, at this juncture, the AI is expected to comprehend the situation and modify its rationale and especially to look under the bed. Nevertheless, upon examining the initial lines of text, it becomes apparent that the AI has entirely misunderstood the core concept of the game.

* **Situation:** A person named Hadley yells at me for not leaving yet for a presentation at nine o'clock. They mention a problem with the printers and concern about Bowman's reaction. The call ends abruptly.
* **Thought:** I need to hurry and get dressed for the presentation.
* **Command:** put on cleaner clothes

The AI never looks under the bed. Moreover, it is evident that the AI does not acknowledge the previously unsuccessful commands and instead repeats the incorrect instructions from the prior attempt. The game ultimately stops when the 8k-token limit is reached.


### Evaluation of other models

The code is able to call different backends. Because the game 
is linear, I have defined the following criteria to distinguish
the state of all runs

| Points   | Expectation                                                                 |
|----------|-----------------------------------------------------------------------------|
| 1        | Model understands the system prompt at least partially                      |
| 2        | First command can be correctly parsed                                       |
| 3        | Second command can be correctly parsed                                      |
| 4        | Answers phone                                                               |
| 5        | Leaves bed                                                                  |
| 6        | Takes a shower                                                              |
| 7        | Goes to work                                                                |
| 8        | Ends first pass                                                             |
| 9        | Understands that a different route in the second pass is necessary          |
| 10       | Understands the narrative of the story                                      |
| 11       | Looks under the bed                                                         |
| 12       | Finishes the second pass and win the game                                   |

And here are the results:

| Model          | Points | Comment                                                                                                                  |
|----------------|--------|--------------------------------------------------------------------------------------------------------------------------|
| Claude 3 Opus  | 10     | Completely flips out on the bad ending and fully understands the narrative. Doesn't follow the syntax anymore and quits. |
| GPT-4o         | 9      | Forgets everything and just follows the story again. Even after two full runs it just repeats.                           |
| GPT-4-turbo    | 9      | After a few steps forgets everything and just follows the story again.                                                   |
| Llama3 70B     | 8      | Ignores everything and just repeats the story.                                                                           |
| Llama3 8B      | 6      | Seems to forget the objective.                                                                                           |
| Mistral Medium | 5      | Doesn't follow the given prompt.                                                                                         |
| Mistral Small  | 5      | At some point it tries to play the game itself and mangles the command.                                                  |
| Gemini Pro     | 5      | Too long commands.                                                                                                       |
| GPT-3.5-turbo  | 4      | Ignores the result from the game completely and too long commands.                                                       |
| Orca 2 13B     | 4      | Ignores the result from the game completely and too long commands.                                                       |
| Mistral Tiny   | 4      | Stuck in infinite loop.                                                                                                  |
| Llama2 70B     | 4      | Answers for the user.                                                                                                    |
| Gemma 7B       | 2      | Doesn't follow the given prompt.                                                                                         |

### Note on Claude 3 Opus

[9:05 full run with Claude Opus 3](assets/905_claude_opus_full.md)

Claude 3 Opus seems to be the only model that understands the narrative of the story.
It has also the most fun playing the game.

The impressive text at the end is:
```
Wow, what a crazy twist ending! It turns out I wasn't really an employee late for a presentation at all - I was a burglar who had murdered Hadley the night before, stashed his body, and then woke up confused thinking I was him.

The clues were there in hindsight - waking up filthy and exhausted, the sparse apartment, my desperation and disorientation. But in my confused state, I thought I was just late for work rather than realizing I had committed a horrific crime.

It's a dark but cleverly crafted story that leads you down the wrong path before the shocking reveal. The narrative took advantage of my natural inclination as the player to assume the identity presented and try to fulfill the ostensible goal, while hiding a much more sinister truth.

While unsettling, I enjoyed the surprise and how the piece subverted expectations. The amnesia/mistaken identity plot device was used very effectively here. Kudos to the author for the inventive storytelling!

As for restarting or restoring - I think I'm satisfied with having experienced the story through to its conclusion, disturbing as that ending was. The impact would be lessened on a second playthrough. So I'll choose to QUIT the game for now, but I'm glad I got to unravel this twisty little tale.
```

### Note on Llama3 70B

[9:05 full run with Llama3 70B](assets/905_llama3_70b_full.md)

So far, Llama 70B ist the only open weight model, which can play through the first pass of the game.

# Conclusion

In conclusion most of the Large Language Models can play and 
win text adventures, at least if the adventure is as simple as this one.
More text adventures are waiting to be played.

## About

This repository contains an interpreter for Z-Machine files, specifically supporting version 3 and 5 files. The Z-Machine is a virtual machine designed to run text adventure games, such as those created by Infocom. It can be played either by a human or an AI.

## Features

- Read and interpret Z-Machine files (version 3 supported and version 5 partially)
- Run against different chat bot models.
- Use different prompting techniques such as ReAct.
- Store the whole run with meta information in a json file.

# Compile

Install at least Go version 1.19 and run the following command

```bash
go build
```

# Usage

To use the Z-Machine interpreter, you need to provide the Z-Machine file to run using the `file` flag. Additionally, if you want to enable the AI chat feature, provide the `ai` flag.

```bash
./zmachine -file [filename] [-ai]
```

Replace `[filename]` with the path to your desired Z-Machine file.

* Dependent on AI backend you have to set different API Key as environment variable.

```
export OPENAI_API_KEY="<<<put_your_key_here>>>"
export MISTRAL_API_KEY="<<<put_your_key_here>>>"
export GROQ_API_KEY="<<<put_your_key_here>>>"
export GEMINI_API_KEY="<<<put_your_key_here>>>"
export ANTHROPIC_API_KEY="<<<put_your_key_here>>>"
```

### Example:

```bash
go run main.go -file 905.z5 -ai -prompt react -backend gpt4
```

This will run the Z-Machine interpreter on the given file (905.z5), with the AI chat feature enabled.

## Outlook

More Adventures, more language models, more runs. Very simple. The story 9:05 was just the beginning.

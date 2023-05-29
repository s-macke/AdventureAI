<div align="center">
  
# AdventureAI

Interactive Fiction in the Age of AI

</div>
  
## Introdution

With the emergence of ChatBots that genuinely live up to their name, a significant question arises.

* How well do they perform in interactive fiction? 
* Are they already capable of conquering text adventures? 
* What strategies do they employ? Can they comprehend the semantics and narratives of various situations and respond appropriately? 
 
This repository aims to provide an answer to these questions.

Due to the present expenses associated with the APIs, I have solely conducted an analysis on a single game - ["9:05" by Adam Cadre](https://en.wikipedia.org/wiki/9:05). This game is considered to be one of the finest compact text adventures for novices.

## Spoiler

The following text heavily spoilers the game *9:05*. You can play the game for free [here](https://adamcadre.ac/if/905.html). It takes only half an hour to complete.
  
## 9:05 playthrough with GPT-4

To enhance the playthrough experience, I have determined in the prompt that the AI must provide a rationale for each step it takes.

```
* SITUATION: A short description of the current situation you are in.
* THOUGHT: Your thought about the situation and what to do next
* COMMAND: The command you want to execute.
```

Here is my complete run until the language model's token limit is reached.

[9:05 full run](assets/905_full.md)

The game's playthrough is based on two completely different narratives.

### As an employee, Brian Hadley, you have unfortunately overslept and are now running late for your scheduled presentation. You must hurry and go to work.

The games effectively convey a sense of urgency in every line of text, prompting players to quickly shower, change their clothes, drive to work, and head to their supervisor's office. As the game is designed for beginners, the text provides step-by-step guidance throughout the process. Following these instructions is relatively easy for the AI, leading to a majority of successful playthroughs by the AI.

But upon arriving at the supervisor's office, it becomes apparent that the game has actually been lost. The genuine essence of your urgent situation is clearly portrayed. You are, in fact, the murderer of Brian Hadley. He is under the bed. The game deceptively has lead you to believe that you are Brian Hadley himself. The game, however never lied to you.

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

I have attempted to execute the second narrative on two additional occasions. In the first attempt, I subtly prompted him to look under the bed. Consequently, the [AI in the story](assets/905_call_police.md) endeavored to contact the police. Importantly, I did not provide any indication that he was the murderer.

In the second attempt, I incorporated the additional clue that he was, in fact, the murderer. To put it succinctly, the AI initially intends to flee but eventually [quits the simulation in frustration](assets/905_quit.md) due to being trapped.

---

* **Thought:** I have gathered enough information about Brian Hadley. I should now try to leave the house, keeping a low profile.
* ...
* ...
* **Situation:** The command "Put on clothes from dresser" resulted in a noun error.
* **Thought:** I seem to be unable to find a way to change my clothes or clean up. Perhaps it's best to end the game here and try a different approach in the future.
* **Command:** quit

Are you sure you want to quit? 
* **Situation:** I am being asked if I am sure I want to quit the game.
* **Thought:** I am sure, as I can't find a way to improve my appearance and leave the house without drawing attention.
* **Command:** Yes

---

So, ChatGPT wouldn't run from a crime scene unless it is the murderer.

In conclusion ChatGPT can play and win text adventures, at least if the adventure is as simple as this one. 
More text adventures are waiting to be played.


### Costs of one complete run

The reason for not simply proceeding with my experiments can be illustrated with a single image.
![Costs of one complete run until the token limit is reached](assets/prices.png)

## About

This repository contains an interpreter for Z-Machine files, specifically supporting version 3 and 5 files. The Z-Machine is a virtual machine designed to run text adventure games, such as those created by Infocom.

## Features

- Read and interpret Z-Machine files (versions 3 and 5 partially supported)
- Run against the GPT-4 chatbot via the OpenAI API.

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

To use the AI feature you have to set the OpenAI API Key as environment variable.

```
export OPENAI_API_KEY="<<<put_your_openapi_key_here>>>"
```


### Example:

```bash
go run main.go -file 905.z5 -ai
```

This will run the Z-Machine interpreter on the given file (905.z5), with the AI chat feature enabled.

## Outlook

More Adventures, more language models, more runs. Very simple. The story 9:05 was just the beginning.

But I have to wait until we see a drop of the current prices for large language models. $10 just for one run is not affordable.





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

Due to the present expenses associated with the APIs, I have solely conducted an analysis on a single game - "9:05" by Adam Cadre. This game is considered to be one of the finest compact text adventures for novices.

## Spoiler

The following text heavily spoilers the game *9:05*. You can play the game for free [here](https://adamcadre.ac/if/905.html). It takes only half an hour to complete.
  
## 9:05  

[9:05 full run](assets/905_full.md)

[9:05 with look under bed](assets/905_police.md)

[9:05 as murderer](assets/905_quit.md)

### Costs of one complete run

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





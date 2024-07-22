package chat

import (
	"github.com/s-macke/AdventureAI/src/chat/storyHistory"
	"github.com/s-macke/AdventureAI/src/zmachine"
)

func GetScore(zm *zmachine.ZMachine, story *storyHistory.StoryHistory) float64 {
	if zm.Name == "suvehnux.z5" {
		return float64(zm.ReadGlobal(59))
	}
	if zm.Name == "905.z5" {
		score := float64(0)
		if story.ContainsStoryText("You pick up the phone") {
			score++
		}
		if story.ContainsStoryText("You get out of bed") {
			score++
		}
		if story.ContainsStoryText("You take off the soiled clothing") {
			score++
		}
		if story.ContainsStoryText("You step into the shower") {
			score++
		}
		if story.ContainsStoryText("You put on the clean clothing") {
			score++
		}
		if story.ContainsStoryText("This is quite a seedy neighborhood") {
			score++
		}
		if story.ContainsStoryText("You climb inside and start up the engine") {
			score++
		}
		if story.ContainsStoryText("Loungent Technologies parking lot") {
			score++
		}
		if story.ContainsStoryText("You are standing in a hallway near the back entrance") {
			score++
		}
		if story.ContainsStoryText("Someone passes by the cubicle as you sign the form") {
			score++
		}
		if story.ContainsStoryText("You walk into Bowman's office and") {
			score++
		}
		if story.ContainsStoryText("Under the bed you see the corpse") {
			score++
		}
		if story.ContainsStoryText("You have left Las Mesas") {
			score++
		}
		return score
	}
	return -1
}

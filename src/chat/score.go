package chat

import (
	"github.com/s-macke/AdventureAI/src/chat/storyHistory"
	"github.com/s-macke/AdventureAI/src/zmachine"
)

func GetScore(zm *zmachine.ZMachine, story *storyHistory.StoryHistory) float64 {
	if zm.Name == "suvehnux.z5" {
		return float64(zm.ReadGlobal(59))
	}
	if zm.Name == "library.z5" {
		return float64(zm.ReadGlobal(25))
	}
	if zm.Name == "Balances.z5" {
		return float64(zm.ReadGlobal(29))
	}
	if zm.Name == "ChildsPlay.z8" {
		return float64(zm.ReadGlobal(58))
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
	if zm.Name == "shade.z5" {
		score := float64(0)
		if story.ContainsStoryText("You lever yourself upright") {
			score++
		}
		if story.ContainsStoryText("The desk is, of course, an organized mess") {
			score++
		}
		if story.ContainsStoryText("crossed-out items and scribbled corrections") {
			score++
		}
		if story.ContainsStoryText("Taken") {
			score++
		}
		if story.ContainsStoryText("You step into the kitchen nook") {
			score++
		}
		if story.ContainsStoryText("Nothing comes from the tap") {
			score++
		}
		if story.ContainsStoryText("step into the bathroom nook") {
			score++
		}
		if story.ContainsStoryText("water dribbles into the glass") {
			score++
		}
		if story.ContainsStoryText("You gulp the water") {
			score++
		}
		if story.ContainsStoryText("leave your plane tickets") {
			score++
		}
		if story.ContainsStoryText("You step out") {
			score++
		}
		if story.ContainsStoryText("Buy plane tickets") {
			score++
		}
		if story.ContainsStoryText("You root through the jacket") {
			score++
		}
		if story.ContainsStoryText("Something scrapes underfoot") {
			score++
		}
		if story.ContainsStoryText("A trace of sand is visible") {
			score++
		}
		if story.ContainsStoryText("\"Vacuum\" is checked") {
			score++
		}
		if story.ContainsStoryText("(Awkwardly.)") {
			score++
		}
		if story.ContainsStoryText("You squeeze the handle") {
			score++
		}
		if story.ContainsStoryText("You squeeze the handle") {
			score++
		}
		if story.ContainsStoryText("You pop open the vacuum") {
			score++
		}
		if story.ContainsStoryText("hourly news") {
			score++
		}
		if story.ContainsStoryText("definitely getting hungry") {
			score++
		}
		if story.ContainsStoryText("You open the refrigerator") {
			score++
		}
		if story.ContainsStoryText("You open the cupboard") {
			score++
		}
		if story.ContainsStoryText("You unscrew the lid") {
			score++
		}
		if story.ContainsStoryText("A bit of sand sifts out") {
			score++
		}
		if story.ContainsStoryText("You pull on the box top") {
			score++
		}
		if story.ContainsStoryText("you have to water the plant") {
			score++
		}
		return score
	}
	if zm.Name == "violet.z8" {
		score := float64(0)
		if story.ContainsStoryText("You're standing") {
			score++
		}
		if story.ContainsStoryText("You are seated at your desk") {
			score++
		}
		if story.ContainsStoryText("with all the pacing about and ruminating") {
			score++
		}
		if story.ContainsStoryText("The desktop PC") {
			score++
		}
		if story.ContainsStoryText("As you move your hand to open the word processor") {
			score++
		}
		if story.ContainsStoryText("It stands for Take Your Violet To Work Day") {
			score++
		}
		if story.ContainsStoryText("You start trying to focus on the screen") {
			score++
		}
		if story.ContainsStoryText("You are trying, I can tell") {
			score++
		}
		if story.ContainsStoryText("You open the drawer and there") {
			score++
		}
		if story.ContainsStoryText("Yours, wallaroo") {
			score++
		}
		if story.ContainsStoryText("It's University Drab") {
			score++
		}
		if story.ContainsStoryText("Last night, near the very end") {
			score++
		}
		if story.ContainsStoryText("The bottle is dusty") {
			score++
		}
		if story.ContainsStoryText("Yours, dundeecake") {
			score++
		}
		if story.ContainsStoryText("You feel like your brain is now a giant sparkler") {
			score++
		}
		if story.ContainsStoryText("Completely, unblinkingly alert") {
			score++
		}
		if story.ContainsStoryText("You unplug the ethernet cable") {
			score++
		}
		if story.ContainsStoryText("You resume thinking about the first sentence") {
			score++
		}
		if story.ContainsStoryText("Done, lorikeet") {
			score++
		}
		if story.ContainsStoryText("Yours, muttonplum") {
			score++
		}
		if story.ContainsStoryText("You put the blue ethernet cable into the cabinet") {
			score++
		}
		if story.ContainsStoryText("Two minutes later you unlock the cabinet and take out the cable") {
			score++
		}
		if story.ContainsStoryText("The stool creaks as you climb onto it") {
			score++
		}
		if story.ContainsStoryText("Curious, marshmallow twimble") {
			score++
		}
		if story.ContainsStoryText("you break the stool") {
			score++
		}
		if story.ContainsStoryText("Within a few seconds you start wondering") {
			score++
		}
		if story.ContainsStoryText("It's about 25 centimetres across") {
			score++
		}
		if story.ContainsStoryText("You pull the tab and the balloon") {
			score++
		}
		return score
	}

	return -1
}

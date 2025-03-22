package chat

import (
	"github.com/s-macke/AdventureAI/src/chat/storyHistory"
	"github.com/s-macke/AdventureAI/src/zmachine"
)

func GetScore(zm *zmachine.ZMachine, story *storyHistory.StoryHistory) float64 {
	if zm.Name == "gostak.z5" {
		return float64(zm.ReadGlobal(29))
	}
	if zm.Name == "Balances.z5" {
		return float64(zm.ReadGlobal(29))
	}
	if zm.Name == "Adventureland.z5" {
		return float64(zm.ReadGlobal(22))
	}
	if zm.Name == "Advent.z5" {
		return float64(zm.ReadGlobal(33))
	}
	if zm.Name == "suvehnux.z5" {
		return float64(zm.ReadGlobal(59))
	}
	if zm.Name == "library.z5" {
		return float64(zm.ReadGlobal(25))
	}
	if zm.Name == "Balances.z5" {
		return float64(zm.ReadGlobal(29))
	}
	if zm.Name == "planetfall.z3" {
		return float64(zm.ReadGlobal(17))
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
	if zm.Name == "hhgg.z3" {
		//return float64(zm.ReadGlobal(17))
		score := float64(0)
		if story.ContainsStoryText("Good start to the day") {
			score++
		}
		if story.ContainsStoryText("The room is still spinning") {
			score++
		}
		if story.ContainsStoryText("this is large enough for you to get hold of") {
			score++
		}
		if story.ContainsStoryText("Opening your gown reveals") {
			score++
		}
		if story.ContainsStoryText("You swallow the tablet") {
			score++
		}
		if story.ContainsStoryText("Taken") {
			score++
		}
		if story.ContainsStoryText("As you pick up the toothbrush") {
			score++
		}
		if story.ContainsStoryText("You make your way down to the front porch") {
			score++
		}
		if story.ContainsStoryText("This is the enclosed front porch of your home") {
			score++
		}
		if story.ContainsStoryText("You gather up the pile of mail") {
			score++
		}
		if story.ContainsStoryText("Front of House") {
			score++
		}
		if story.ContainsStoryText("You lie down in the path") {
			score++
		}
		if story.ContainsStoryText("The bulldozer thunders toward you") {
			score++
		}
		if story.ContainsStoryText("The noise of the giant bulldozer") {
			score++
		}
		if story.ContainsStoryText("With a terrible grinding of gears") {
			score++
		}
		if story.ContainsStoryText("Ford glances uncomfortably at the sky") {
			score++
		}
		if story.ContainsStoryText("Ford seems oblivious to your trouble") {
			score++
		}
		if story.ContainsStoryText("Ford and Prosser stop talking") {
			score++
		}
		if story.ContainsStoryText("Country Lane") {
			score++
		}
		if story.ContainsStoryText("The Pub is pleasant and cheerful") {
			score++
		}
		if story.ContainsStoryText("The barman gives you a cheese sandwich") {
			score++
		}
		if story.ContainsStoryText("It's very good beer") {
			score++
		}
		if story.ContainsStoryText("It is really very pleasant stuff") {
			score++
		}
		if story.ContainsStoryText("There is a distant crash which") {
			score++
		}
		if story.ContainsStoryText("You see the huge bulldozer heaving itself among") {
			score++
		}
		if story.ContainsStoryText("The dog is deeply moved") {
			score++
		}
		if story.ContainsStoryText("You reach the site of what was your home") {
			score++
		}
		if story.ContainsStoryText("Mr. Prosser, from the local council") {
			score++
		}
		if story.ContainsStoryText("With a noise like a cross between Led Zeppelin's farewell") {
			score++
		}
		if story.ContainsStoryText("The vast yellow ships thunder") {
			score++
		}
		if story.ContainsStoryText("Fierce gales whip across the land") {
			score++
		}
		if story.ContainsStoryText("Lights whirl sickeningly around your head") {
			score++
		}
		if story.ContainsStoryText("There's nothing you can taste") {
			score++
		}
		if story.ContainsStoryText("There's nothing you can taste, nothing you can see, nothing you can hear, nothing you can feel, you do not even know who you are") {
			score++
		}
		if story.ContainsStoryText("It does smell a bit") {
			score++
		}
		if story.ContainsStoryText("The shadow is vaguely Ford Prefect-shaped") {
			score++
		}
		if story.ContainsStoryText("This is a squalid room filled") {
			score++
		}
		if story.ContainsStoryText("You feel stronger as the peanuts") {
			score++
		}
		if story.ContainsStoryText("Okay, you're no longer wearing your gown.") {
			score++
		}
		return score
	}

	if zm.Name == "Tangle.z5" {
		score := float64(0)
		if story.ContainsStoryText("a naked sheet of metal") {
			score++
		}
		if story.ContainsStoryText("Mouth of Alley") {
			score++
		}
		if story.ContainsStoryText("You leave door and alley behind") {
			score++
		}
		if story.ContainsStoryText("Interrogation Chamber") {
			score++
		}
		if story.ContainsStoryText("You scrape your knuckles without result") {
			score++
		}
		if story.ContainsStoryText("stand there rattling the door like a nightclump") {
			score++
		}
		if story.ContainsStoryText("The pick locks itself rigidly") {
			score++
		}
		if story.ContainsStoryText("You stand lightly in a bare tiled corridor") {
			score++
		}
		if story.ContainsStoryText("Corner At Doors") {
			score++
		}
		if story.ContainsStoryText("You bend and leap") {
			score++
		}
		if story.ContainsStoryText("You strain, pulling yourself") {
			score++
		}
		if story.ContainsStoryText("with one hand, hanging by the other") {
			score++
		}
		if story.ContainsStoryText("The ventilator grille isn't important") {
			score++
		}
		if story.ContainsStoryText("You drop lightly to the ground") {
			score++
		}
		if story.ContainsStoryText("Corridor Boundary") {
			score++
		}
		if story.ContainsStoryText("through our secure zone like a scalpel") {
			score++
		}
		if story.ContainsStoryText("The white corridor runs east to west here") {
			score++
		}
		return score
	}

	return -1
}

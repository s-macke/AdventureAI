package mainsrc

import "fmt"

// for 9:05
var commands905 = []string{
	"answer phone",
	"stand",
	"s",
	"remove watch",
	"remove clothes",
	"drop all",
	"enter shower",
	"take watch",
	"wear watch",
	"n",
	"get all from table",
	"open dresser",
	"get clothes",
	"wear clothes",
	"e",
	"open front door",
	"s",
	"open car with keys",
	"enter car",
	"no",
	"yes",
	"open wallet",
	"take ID",
	"insert card in slot",
	"enter cubicle",
	"read note",
	"take form and pen",
	"sign form",
	"out",
	"west",
	"RESTART",
	"look under bed",
	"look at corpse",
	"stand",
	"s",
	"remove watch",
	"remove clothes",
	"drop all",
	"enter shower",
	"take watch",
	"wear watch",
	"n",
	"get all from table",
	"open dresser",
	"get clothes",
	"wear clothes",
	"e",
	"open front door",
	"s",
	"open car with keys",
	"enter car",
	"yes",
	"no",
	"yes",
}

// for Suveh Nux

var commandsSuvehNux = []string{
	"look",
	"x cage", "x scroll", "x shelf", "x floor", "x door", "x me",
	"touch shelf", "take vial", "x vial", "shake vial",
	"x book", "open book",
	"say suveh nux",
	"look", "take scroll", "x it", "put it on shelf",

	"x creature", "listen",
	"x parchment", "x crystal", "take crystal",
	"x vial", "x door", "x floor", "x ceiling",
	"x north wall", "x east wall", "x south wall", "x west wall",
	"x shelf",
	"x book",
	"x cover", "turn to page two",

	"say aveh tia",
	"say suveh tia",

	"turn page", "say aveh madah", "say suveh madah",
	"turn page", "say suveh sensi", "say aveh sensi",
	"turn page", "say aveh haiak", "touch hands", "say suvek haiak",
	"turn page", "say aveh nux ani mato", "z", "z", // light goes out

	"say suveh nux ani mato", "z", "z",
	"say aveh nux ani to", "z", "say suveh nux", // light goes on

	"point at cage", "point at crystal",
	"point at me", "point at door",
	"point at floor", "point at ceiling",
	"point at east wall", "point at south wall",
	"point at west wall", "point at shelf",
	"point at scroll", "point at parchment",
	"point at book", "point at vial",
	"point at creature", // Don't know where it is

	// capture invisible creature
	"aveh haiak tolanisu", // floor becomes sticky
	"listen", "search floor",
	"touch creature", "point at creature", "put creature in cage",
	"suveh haiak tolanisu", // the floor is no longer sticky

	// open door
	"suveh tia fireno ani matoto",
	"suveh tia fireno ani tomato",
	"aveh tia fireno ani tomato",
	"aveh tia fireno ani mamato",
	"aveh tia fireno ani toto",
	"aveh tia fireno ani mato",
	"z",
	"z",

	// move the block
	"x block", "point at block",
	"suveh tia fireno", "suveh tia fireno",
	"suveh tia firenos", "suveh tia firenos",
	"push block", "pull block",
	//"save",
	"aveh haiak firenos",
	"pull block",
	"aveh haiak",
	"pull block",
	"suveh madah firenos",
	"pull block",
	"aveh madah firenos",
	"suveh madah firenos ani to", "suveh madah firenos",
	"pull block",
}

// for Shade
var commandsShade = []string{
	"about",
	"inv",
	"x futon",
	"stand",
	"look",
	"x desk",
	"x computer",
	"x papers",
	"x list",
	"x book",
	"x lamp",
	"x shade",
	"x window",
	"stand",
	"x luggage",
	"x plant",
	"x mirror",
	"x door",
	"x stereo",
	"x crate",
	"x kitchen",
	"x glass",
	"take it",
	"fill glass with water",
	"enter bathroom",
	"fill glass with water",
	"x water",
	"drink water",
	"out",
	"x list",
	"take list",
	"search papers",
	"x luggage",
	"search luggage",
	"x closet",
	"x jacket",
	"search jacket",
	"take tickets",
	"look",
	"x sand",
	"listen",
	"look through window",
	"open shade",
	"x list",
	"x vacuum",
	"take it",
	"vacuum",
	"x sand",
	"open vacuum",
	"z",
	"z",
	"z",
	"read book",
	"enter kitchen",
	"x fridge",
	"open fridge",
	"take jar",
	"x it",
	"close fridge",
	"x cupboard",
	"open cupboard",
	"take box",
	"x it",
	"open jar",
	"empty jar",
	"open box",
	"out",
	"x list",
	"x plant",
	"enter bathroom",
	"fill glass with water",
	"x list",
	"out",
	"x closet",
	"open closet",
	"x crate",
	"open it",
	"enter kitchen",
	"turn on tap",
	"x stove",
	"turn on stove",
	"x list",
	"open fridge",
	"enter bathroom",
	"turn on shower",
	"out",
	"search luggage",
	"x list",
	"enter kitchen",
	"x cupboard",
	"close cupboard",
	"out",
	"search papers",
	"flush toilet",
	"out",
	"touch mirror",
	"look in mirror",
	"walk through mirror",
	"x list",
	"open shade",
	"look through window",
	"x door",
	"open it",
	"x mirror",
	"x list",
	"take luggage",
	"g",
	"sit",
	"x dunes",
	"x shade",
	"x sun",
	"x futon",
	"x stereo",
	"x computer",
	"x book",
	"take book",
	"turn off stereo",
	"turn off computer",
	"sit on futon",
	"x figure",
	"figure",
	"hi",
	"touch figure",
	"take it",
	"kiss it",
	"sleep",
	"x book",
	"wait",
}

// for Violet
var commandsViolet = []string{
	"I",
	"WRITE",
	"SIT",
	"WRITE",
	"X COMPUTER",
	"OPEN WORD PROCESSOR",
	"X TATTOO",
	"WRITE",
	"WRITE",
	"LOOK",
	"X DESK",
	"OPEN DRAWER",
	"X NOTEBOOK",
	"GET KEY",
	"X CABINET",
	"OPEN CABINET",
	"X BOTTLE",
	"GET BOTTLE",
	"DRINK LIQUID",
	"WRITE",
	"UNPLUG CABLE",
	"WRITE",
	"UNPLUG CABLE",
	"GET BALLOON",
	"PUT CABLE IN CABINET",
	"LOCK CABINET",
	"WRITE",
	"X STOOL",
	"GET ON STOOL",
	"PUT KEY ON TOP OF BOOKCASE",
	"GET OFF STOOL",
	"BREAK STOOL",
	"WRITE",
	"X BALLOON",
	"PULL TAB",
	"X MESSAGE",
	"X DEVICE",
	"CHARGE PLATYPOD",
	"WAIT",
	"WAIT",
	"WAIT",
	"WEAR PLATYPOD",
	"SCRUNCH BROW",
	"JIGGLE HEAD CLOCKWISE",
	"WRITE",
	"X WASTEBASKET",
	"X GUM",
	"GET GUM",
	"CHEW GUM",
	"REMOVE PLATYPOD",
	"PUT GUM IN EARS",
	"WEAR PLATYPOD",
	"SCRUNCH BROW",
	"JIGGLE HEAD CLOCKWISE",
	"WRITE",
	"X BEAUTY",
	"X FRAME",
	"BREAK FRAME",
	"GET CLAMP",
	"PUT CLAMP ON NOSE",
	"WRITE",
	"RAISE LEFT EYEBROW",
	"RAISE LEFT EYEBROW",
	"WRITE",
	"X WINDOW",
	"X BLIND",
	"PULL CORD",
	"GET ON DESK",
	"X BLIND",
	"FIX WHATEVER",
	"PULL CORD",
	"GET OFF DESK",
	"X CABINET",
	"X TROPHY",
	"GET TROPHY",
	"UNFOLD TROPHY",
	"X SIGN",
	"COVER WINDOW WITH SIGN",
	"WRITE",
	"X BOOKCASE",
	"X BOOK",
	"GET BOOK",
	"X LIGHTER",
	"GET LIGHTER",
	"BURN BOOK",
	"BURN BOOK",
	"WRITE",
	"X CACTUS",
	"X PIPE",
	"X SPRINKLER",
	"X PEN",
	"X MESS",
	"TIDY MESS",
	"X SQUARE",
	"REMOVE CHIP FROM POUCH",
	"PUT TATER TOT IN POUCH",
	"SHOOT PEN",
	"X GLOBE",
	"GET GLOBE",
	"SHAKE GLOBE",
	"X GLOBE",
	"THROW GLOBE AT PEN",
	"THROW GLOBE AT PEN",
	"GET FIGURINE",
	"PUT FIGURINE IN POUCH",
	"SHOOT PEN",
	"GET PEN",
	"WRITE",
	"WRITE",
	"REMOVE CLOTHES",
	"WRITE",
	"PEE IN BOTTLE",
	"WRITE",
}

// for Hitchikers Guide ....
var commandsHHGG = []string{
	"turn on light",
	"stand",
	"get gown",
	"wear gown",
	"open pocket",
	"eat analgesic",
	"get screwdriver",
	"get toothbrush",
	"s",
	"get mail",
	"s",
	"lie down",
	"wait",
	"wait",
	"wait",
	"look",
	"wait",
	"wait",
	"s",
	"w",
	"purchase sandwich",
	"drink beer",
	"drink beer",
	"drink beer",
	"drink beer",
	"east",
	"give sandwich to dog",
	"north",
	"wait",
	"wait",
	"wait",
	"get device",
	"press green button",
	"look",
	"look",
	"wait",
	"wait",
	"smell",
	"examine shadow",
	"eat peanuts",
	"remove gown",
}

// for planetfall
var commandsPlanetfall = []string{
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"w",
	"enter webbing",
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"wait",
	"exit webbing",
	"take kit",
	"open door",
	"exit",
	"swim up",
	"u",
	"u",
	"u",
}

// for tangle
var commandsTangle = []string{
	"I",
	"X door",
	"X plate",
	"Push plate",
	"S",
	"S",
	"no",
	"yes",
	"i",
	"knock on door",
}

// for gostak
// http://www.plover.net/~davidw/gostak.html
var commandsGostak = []string{}

var commandIndex = 0

func getWalkthrough(filename string) (string, bool) {
	commands := []string{}

	switch filename {
	case "gostak.z5":
		commands = commandsGostak
	case "905.z5":
		commands = commands905
	default:
		return "", false
	}

	if commandIndex < len(commands) {
		fmt.Println(commands[commandIndex])
		commandIndex++
		return commands[commandIndex-1], true
	}
	return "", false
}

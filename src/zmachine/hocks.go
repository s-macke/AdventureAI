package zmachine

import "strings"

func HookCommand(zm *ZMachine, input string) bool {
	input = strings.ToLower(input)
	if zm.Name == "gostak.z5" {
		if strings.HasPrefix(input, "jallon") {
			zm.Output.WriteString("P = plazzy fesh\n")
			zm.Output.WriteString("RETURN = tunk fesh\n")
			zm.Output.WriteString("RETURN = tunk fesh\n\n")

			zm.Output.WriteString("### Bewly dedges ###\n\n")
			zm.Output.WriteString("This is a halpock. As in all halpocks, you doatch at it about what to do in a camling by louking murr English dedges, and the halpock louks back. Durly, you could rask a chuld (if you rebbed one) by louking:\n\n  >RASK THE CHULD\n  Rasked.\n\nYou could then tunk the chuld, or pob it at a tendo, or whatever:\n\n  >DURCH CHULD\n  Nothing heamy results.\n\n  >POB MY CHULD AT THE GHARMY TENDO\n  The tendo leils it.\n\nYou can use multiple objects with some dapes:\n\n  >RASK THE GELN, THE POTCHNER, AND THE FACK\n  >POB ALL EXCEPT DRAGUE\n\nYou can also louk multiple dedges by using \".\" or THEN:\n\n  >RASK GELN THEN RASK POTCHNER. POB FACK.\n\nOne very heamy dape is TUNK. You'll be using TUNK a lot, so you can disengope it to T:\n\n  >T TENDO\n  The tendo isn't gharmy now that it's leiled a chuld.\n\nThe halpock recognises many other dapes! Try anything. If the halpock louks a dape, maybe you can louk it back.\n")

			zm.Output.WriteString("### Pelling ###\n\n")
			zm.Output.WriteString("To pell from deave to deave, just louk with that fesh! You must also louk the lutt to pell at.\n\n  >PELL AT THE LOFF LUTT\n  >PELL AT LOFF\n\nYou can also louk this less gopely:\n\n  >PELL L\n\n...or even:\n\n  >L\n\nLikewise, HOFF can be disengoped to H, JIHOFF to JH, and so on.\n\nYou can also pell at deaves. If you rebbed a mosteg, you could:\n\n  >PELL AT THE MOSTEG\n  >GOMB THE MOSTEG\n\nOf course, you can't pell at all lutts from all deaves. And even if you could, there are glauds...")

			zm.Output.WriteString("### Doatching ###\n\n")
			zm.Output.WriteString("Doatching is sloagy in this halpock, and a gope tunder of brolges. If you reb someone to doatch at, louk:\n\n  >DOATCH ABOUT something AT someone\n\nor:\n\n  >DOATCH AT someone ABOUT something\n\nThe fesh of your doatching can be anything you've rebbed or anything someone has doatched at you about.\n\nAs in other halpocks, you can also doatch dedges for others by using the fetticle (\",\").  To doatch at a pank to rask a koshle, you would louk:\n\n  >PANK, RASK THE KOSHLE\n\n...but in this halpock, such dedges will never do anything heamy.\n")

			zm.Output.WriteString("### Heamy dapes ###\n\n")
			zm.Output.WriteString("There are some heamy dapes for dedges about the halpock itself:\n\nQUIT: Quits the halpock.\n\nSAVE/RESTORE: These dapes let you regomb the halpock to an osta. SAVE tunds the halpock to your feshary, and RESTORE regombs it.\n\nRESTART: Regombs the halpock as it was before you louked any dedges at all.\n\nUNDO: Undoes the osta camling, as if it had been distunked.\n\nAGAIN (or G): Relouks your dedge from the osta camling.\n\nOOPS (or O): Oopses a mislouk in the osta dedge. Whatever you louk in the dedge with OOPS will be oopsed into the mislouk.\n\nVERBOSE/BRIEF/SUPERBRIEF: Rikes the halpock's louking when you gomb a deave. VERBOSE louks gope tunkage of deaves whenever you gomb them. BRIEF louks gope tunkage of deaves if you haven't been there before, and ungope tunkage\notherwise. SUPERBRIEF always louks ungope tunkage of deaves.\n\nREB (or R): Louks gope tunkage of your deave.\n\nSCORE: Louks your pellage in the halpock.\n\nMILM (or M): Tunks your raskage.\n\nSCRIPT ON/SCRIPT OFF: SCRIPT ON tunds a pashual of the halpock to your feshary or tavidlouker. SCRIPT OFF undoes SCRIPT ON.\n\nZONCHA (or Z): Zonchas for a camling.")
			zm.Output.WriteString("\n\n>")
			return true
		}
	}

	return false
}

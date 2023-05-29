# 9:05 by Adam Cadre

```
You act as a player of an interactive text adventure. The goal is to win the game. 
The user provides the text of the text adventure. He is not a human and just prints the output of the game.

The format of your output must be:
NARRATIVE: A short description of the current narrative you are in
THOUGHT: Your thought about the situation and what to do next
COMMAND: The command you want to execute. Must always begin with a verb. The maximum number of words are 4. The commands should be very simple.
```




The phone rings.

Oh, no - how long have you been asleep? Sure, it was a tough night, but... This is bad. This is very bad.

The phone rings.

  -----
9:05 by Adam Cadre
v1.12 (2016.0430) / Original release 2000.0103
  -----

Bedroom (in bed)
This bedroom is extremely spare, with dirty laundry scattered haphazardly all over the floor. Cleaner clothing can be found in the dresser. A bathroom lies to the south, while a door to the east leads to the living room.

On the end table are a telephone, a wallet and some keys.

The phone rings.

>
> * **Situation:** I am in a messy bedroom with dirty laundry on the floor. The dresser contains cleaner clothes. There is a bathroom to the south and a living room to the east. A telephone, wallet, and keys are on the end table. The phone is ringing.
> * **Thought:** I should answer the phone to find out who is calling.
> * **Command:** answer the telephone

You pick up the phone. "Hadley!" a shrill voice cries. "Hadley, haven't you even left yet?? You knew that our presentation was at nine o'clock sharp! First the thing with the printers, now this - there won't even be enough left of you for Bowman to fire once he's done with you. Now get the hell down here!!"

Click.

>
> * **Situation:** A person named Hadley yells at me for not leaving yet for a presentation at nine o'clock. They mention a problem with the printers and concern about Bowman's reaction. The call ends abruptly.
> * **Thought:** I need to hurry and get dressed for the presentation.
> * **Command:** put on cleaner clothes

[!: Noun error. (This could mean a number of things: a noun used in your command may not be present, or it may be present but not visible; it may be scenery you can't interact with; you may have specified too many indirect objects; you may have misspelled something; the list goes on.)]

>
> * **Situation:** I am unable to pick up the cleaner clothes directly.
> * **Thought:** I should open the dresser to get the cleaner clothes.
> * **Command:** open the dresser

You'll have to get out of bed first.

>
> * **Situation:** I am still in bed and must get out before accessing the dresser.
> * **Thought:** I should get out of bed and then open the dresser.
> * **Command:** get out of bed

You get out of bed.

>
> * **Situation:** I am now out of bed, standing in the bedroom.
> * **Thought:** Now that I'm out of bed, I should open the dresser and get dressed.
> * **Command:** open the dresser

You open the dresser, revealing some clean clothing.

>
> * **Situation:** The dresser is open, and clean clothing is visible inside.
> * **Thought:** I should put on the clean clothes quickly to attend the presentation.
> * **Command:** put on clean clothes

(first taking the clean clothing)
You'll have to take off the soiled clothes you're wearing first.

>
> * **Situation:** I need to remove the soiled clothes I'm currently wearing before putting on the clean ones.
> * **Thought:** I should take off the soiled clothes first.
> * **Command:** take off soiled clothes

You take off the soiled clothing.

>
> * **Situation:** I have removed the soiled clothing and am now undressed.
> * **Thought:** I should put on the clean clothes from the dresser.
> * **Command:** put on clean clothes

It seems ridiculous to put on fresh, clean clothes when you yourself are covered with mud and dried sweat.

>
> * **Situation:** I am covered in mud and dried sweat, making it unreasonable to put on clean clothes without cleaning myself first.
> * **Thought:** I should take a quick shower in the bathroom to the south.
> * **Command:** go to bathroom


Bathroom
This is a far from luxurious but still quite functional bathroom, with a sink, toilet and shower. The bedroom lies to the north.

>
> * **Situation:** I am now in a basic, functional bathroom with a sink, toilet, and shower. The bedroom is to the north.
> * **Thought:** I should take a quick shower to clean myself before getting dressed.
> * **Command:** take a shower

You don't think the watch you're wearing is waterproof.

>
> * **Situation:** I am wearing a watch that might not be waterproof.
> * **Thought:** I should remove my watch before taking a shower.
> * **Command:** remove the watch

You take off the gold watch.

>
> * **Situation:** I have removed the gold watch from my wrist.
> * **Thought:** Now that the watch is off, I can take a shower.
> * **Command:** take a shower

Carrying your possessions into the shower is a bad idea.

>
> * **Situation:** I am carrying possessions that shouldn't be taken into the shower.
> * **Thought:** I should put down my possessions before showering.
> * **Command:** put down possessions

[!: Noun error. (This could mean a number of things: a noun used in your command may not be present, or it may be present but not visible; it may be scenery you can't interact with; you may have specified too many indirect objects; you may have misspelled something; the list goes on.)]

>
> * **Situation:** The command to put down my possessions was not recognized.
> * **Thought:** I need to specify the items I'm carrying.
> * **Command:** list possessions

[!: Verb error. (This generally means that the first word of your command was not recognized.)]

>
> * **Situation:** My command to list my possessions was not recognized.
> * **Thought:** I should try to remember what I'm carrying and put those items down.
> * **Command:** put down watch

Dropped.

>
> * **Situation:** I have put down the gold watch.
> * **Thought:** I should recall the other items I picked up earlier and put them down.
> * **Command:** put down soiled clothes

Dropped.

>
> * **Situation:** I have put down the soiled clothes.
> * **Thought:** Now that I've put down the items, I should take a shower.
> * **Command:** take a shower

Carrying your possessions into the shower is a bad idea.

>
> * **Situation:** I am still carrying possessions that shouldn't be taken into the shower.
> * **Thought:** I must have missed some items; I should try to remember and put them down as well.
> * **Command:** put down keys

[!: Noun error. (This could mean a number of things: a noun used in your command may not be present, or it may be present but not visible; it may be scenery you can't interact with; you may have specified too many indirect objects; you may have misspelled something; the list goes on.)

If you are unfamiliar with the conventions of interactive fiction, you may wish to play through the tutorial at http://adamcadre.ac/if.html - it explains how to phrase commands in a way that produces fewer errors.]

>
> * **Situation:** My command to put down the keys was not recognized.
> * **Thought:** I should make sure I'm not carrying any other items before showering.
> * **Command:** put down wallet

[!: Noun error.]

>
> * **Situation:** My command to put down the wallet was not recognized.
> * **Thought:** Perhaps I'm not carrying those items after all. I should try taking a shower again.
> * **Command:** take a shower

Carrying your possessions into the shower is a bad idea.

>
> * **Situation:** I am still carrying possessions that shouldn't be taken into the shower.
> * **Thought:** I must not have identified all the items I'm carrying. I need to find alternative ways to find out.
> * **Command:** inventory

You are carrying:
  some clean clothing

>
> * **Situation:** I am carrying clean clothing.
> * **Thought:** I should put down the clean clothing before taking a shower.
> * **Command:** put down clean clothing

Dropped.

>
> * **Situation:** I have put down the clean clothing.
> * **Thought:** Now that I've put down all my possessions, I can take a shower.
> * **Command:** take a shower

You step into the shower, turn on the water, and within a few moments you feel like a new man. But no time to dawdle - you hop back out again and dry off in record time.

>
> * **Situation:** I quickly took a shower, feeling refreshed. I wasted no time drying off and getting ready to move on.
> * **Thought:** I should pick up my items, put my clean clothes on, and head to the presentation as quickly as possible.
> * **Command:** put on clean clothes

(first taking the clean clothing)
You put on the clean clothing.

>
> * **Situation:** I have put on the clean clothing.
> * **Thought:** Now that I'm dressed, I should pick up my possessions and head to the living room.
> * **Command:** pick up possessions

[!: Noun error.]

>
> * **Situation:** My command to pick up my possessions was not recognized.
> * **Thought:** I should specify each item and pick them up individually.
> * **Command:** pick up watch

Taken.

>
> * **Situation:** I have picked up the gold watch.
> * **Thought:** I should pick up the rest of my possessions before leaving.
> * **Command:** pick up wallet

[!: Noun error.

Again, if you're getting frustrated, please consider playing through the tutorial at http://adamcadre.ac/if.html.]

>
> * **Situation:** My command to pick up the wallet was not recognized.
> * **Thought:** I should remember to pick up the soiled clothes and any other possessions.
> * **Command:** pick up soiled clothes

Taken.

>
> * **Situation:** I have picked up the soiled clothes.
> * **Thought:** I should continue to pick up any other possessions before leaving.
> * **Command:** pick up keys

[!: Noun error.]

>
> * **Situation:** My command to pick up the keys was not recognized.
> * **Thought:** I may have missed some items earlier. I should go to the living room and search for the wallet and keys.
> * **Command:** go to living room

[!: Noun error.]

>
> * **Situation:** My command to go to the living room was not recognized.
> * **Thought:** I should try another command to exit the bathroom and head to the living room.
> * **Command:** go to bedroom


Bedroom
This bedroom is extremely spare, with dirty laundry scattered haphazardly all over the floor. A bathroom lies to the south, while a door to the east leads to the living room.

On the end table are a telephone, a wallet and some keys.

>
> * **Situation:** I am back in the bedroom, where the floor is covered in dirty laundry. There is a bathroom to the south and a living room to the east. A telephone, wallet, and keys are on the end table.
> * **Thought:** I should pick up the wallet and keys from the end table.
> * **Command:** pick up wallet and keys

wallet: Taken.
keys: Taken.

>
> * **Situation:** I have picked up the wallet and keys from the end table.
> * **Thought:** Now that I have my items, I should head to the living room and leave for the presentation.
> * **Command:** go to living room


Living room
The furniture isn't as carefully arranged as one might like, and it's a bit lacking in the luxuries that might make this a more pleasant place to spend an evening - no stereo, no TV, no DVD, none of that - but this living room is still nicer than most of those in this neighborhood. A kitchenette lies to the east, the bedroom to the west and the front door to the south.

>
> * **Situation:** I am now in the living room, which has basic furniture and lacks luxury items such as a stereo, TV, or DVD. A kitchenette is to the east, the bedroom is to the west, and the front door is to the south.
> * **Thought:** I should head to the front door and leave for the presentation.
> * **Command:** go to front door

The front door is closed.

>
> * **Situation:** The front door is closed.
> * **Thought:** I should open the front door and head outside.
> * **Command:** open front door

You open the front door.

>
> * **Situation:** I have opened the front door.
> * **Thought:** Now that the door is open, I should head outside and go to the presentation.
> * **Command:** go outside


Driveway
This is quite a seedy neighborhood, but there isn't really anywhere in the whole Las Mesas metro area that isn't at least a little rundown. Of course, the locals may complain about the ugliness and the crime, but the fact is that a starter home like this one would cost easily five times as much in someplace like Las Playas - and that's if you could find one.

The house lies to the north. A car is parked here in the driveway.

>
> * **Situation:** I am now outside in the driveway. The neighborhood is seedy and rundown. The house is to the north, and a car is parked in the driveway.
> * **Thought:** I should get into the car and drive to the presentation.
> * **Command:** enter the car

The car is locked.

>
> * **Situation:** The car is locked.
> * **Thought:** I should unlock the car with the keys I picked up earlier.
> * **Command:** unlock car with keys

[!: Command not understood past "unlock the car."]

>
> * **Situation:** My command to unlock the car with keys was not understood completely.
> * **Thought:** I should try a simpler command to unlock the car.
> * **Command:** unlock car

Unlocked.

>
> * **Situation:** I have successfully unlocked the car.
> * **Thought:** Now that the car is unlocked, I should get in and drive to the presentation.
> * **Command:** enter the car

You climb inside and start up the engine.

Driving
Ah, scenic Las Mesas. Man, this place is an absolute toilet. Soon you'll be able to afford to get the hell out of here - provided you can avoid making any more slip-ups on the job.

As you cruise down the road, you notice a freeway onramp approaching. Would you like to get on? >> 
> * **Situation:** I am now inside the car and driving through the unattractive Las Mesas. There is a freeway onramp approaching as I drive.
> * **Thought:** To reach the presentation as quickly as possible, I should take the freeway onramp.
> * **Command:** take freeway onramp

Please answer yes or no.> 
> * **Situation:** The game is asking me to answer yes or no about taking the freeway onramp.
> * **Thought:** I should clarify my response.
> * **Command:** yes


You start to change lanes, but you notice that the ramp is out of service today - some sort of maintenance work. But that's okay - there's another onramp less than a mile down the road from here. You keep driving.

As you cruise down the road, you see the familiar Loungent Technologies building approaching. Would you like to stop here? >> 
> * **Situation:** The initial freeway onramp is out of service, but there's another one further down the road. I am now driving near the Loungent Technologies building.
> * **Thought:** I should stop at the Loungent Technologies building for the presentation.
> * **Command:** yes


You pull into the parking lot and hop out.

Loungent Technologies parking lot
You are at the employees' entrance to the Loungent Technologies building. There is a slot by the door, next to a sign reading "INSERT ID CARD HERE".

The car is parked in the parking lot.

>
> * **Situation:** I have arrived at the Loungent Technologies parking lot and stand near the employee entrance. There is a slot by the door with a sign indicating where to insert an ID card. The car is parked in the parking lot.
> * **Thought:** I should find my ID card and insert it into the slot.
> * **Command:** find ID card

[!: Verb error. (This generally means that the first word of your command was not recognized.)]

>
> * **Situation:** My command to find the ID card was not recognized.
> * **Thought:** I should search my wallet for the ID card.
> * **Command:** search wallet

You can't see inside, since the wallet is closed.

>
> * **Situation:** I can't see inside the wallet because it's closed.
> * **Thought:** I should open the wallet to find the ID card.
> * **Command:** open wallet

You open the wallet, revealing a driver's license and an ID card.

>
> * **Situation:** I have opened the wallet, and there is a driver's license and an ID card inside.
> * **Thought:** I should take the ID card and insert it into the slot by the entrance.
> * **Command:** take ID card

Taken.

>
> * **Situation:** I have taken the ID card from the wallet.
> * **Thought:** Now that I have the ID card, I should insert it into the slot by the entrance to get inside the building.
> * **Command:** insert ID card into slot

A green LED lights up, the door clicks open and you step inside.

Loungent Technologies
You are standing in a hallway near the back entrance to Loungent Technologies. The hallway leads north to the reception area and south to the parking lot, while a door marked "MATTHEW BOWMAN" lies to the west.

There is a cubicle here; it is marked simply "HADLEY".

>
> * **Situation:** I am now inside Loungent Technologies, in a hallway near the back entrance. The hallway leads north to the reception area, south to the parking lot, and a door marked "MATTHEW BOWMAN" is to the west. There's a cubicle marked "HADLEY" nearby.
> * **Thought:** I should quickly go to my cubicle and see if I've missed anything important before the presentation.
> * **Command:** go to Hadley's cubicle

[!: Noun error.]

>
> * **Situation:** My command to go to Hadley's cubicle was not recognized.
> * **Thought:** I should try rephrasing the command to enter the cubicle.
> * **Command:** enter Hadley's cubicle

[!: Noun error.]

>
> * **Situation:** My command to enter Hadley's cubicle was not recognized.
> * **Thought:** I should find another way to approach Hadley's cubicle.
> * **Command:** approach Hadley's cubicle

[!: Verb error. (This generally means that the first word of your command was not recognized.)

If you are unfamiliar with the conventions of interactive fiction, you may wish to play through the tutorial at http://adamcadre.ac/if.html - it explains how to phrase commands in a way that produces fewer errors.]

>
> * **Situation:** My command to approach Hadley's cubicle was not recognized.
> * **Thought:** I'll try one more time using a different command to get to Hadley's cubicle.
> * **Command:** go to cubicle

You get into the cubicle.

On the desk are a form, a pen and a note.

>
> * **Situation:** I am now inside Hadley's cubicle. On the desk, there is a form, a pen, and a note.
> * **Thought:** I should read the note to see if it contains any important information.
> * **Command:** read note

The note reads, "Hadley - sign this F209 and return it to me IMMEDIATELY!! MB".

>
> * **Situation:** The note reads, "Hadley - sign this F209 and return it to me IMMEDIATELY!! MB."
> * **Thought:** I should sign the F209 form and return it to MB as soon as possible.
> * **Command:** sign F209 form

You're not holding a pen.

>
> * **Situation:** I am not holding a pen to sign the F209 form.
> * **Thought:** I should pick up the pen and sign the F209 form.
> * **Command:** pick up pen

Taken.

>
> * **Situation:** I have picked up the pen.
> * **Thought:** Now that I have the pen, I should sign the F209 form.
> * **Command:** sign F209 form

Someone passes by the cubicle as you sign the form. You look up just in time to see Bowman's door click shut.

>
> * **Situation:** I signed the F209 form, and as I did so, someone passed by the cubicle. I looked up and saw Bowman's door closing.
> * **Thought:** I should return the signed F209 form to Bowman as instructed.
> * **Command:** go to Bowman's door

You'll have to get out of the cubicle first.

>
> * **Situation:** I need to get out of the cubicle before heading to Bowman's door.
> * **Thought:** I should exit the cubicle and then approach Bowman's door.
> * **Command:** exit cubicle

You get out of the cubicle.

Loungent Technologies
You are standing in a hallway near the back entrance to Loungent Technologies. The hallway leads north to the reception area and south to the parking lot, while a door marked "MATTHEW BOWMAN" lies to the west.

There is a cubicle here; it is marked simply "HADLEY".

>
> * **Situation:** I am now outside the cubicle, back in the hallway near the back entrance. The hallway leads north to the reception area, south to the parking lot, and a door marked "MATTHEW BOWMAN" is to the west. The cubicle marked "HADLEY" is nearby.
> * **Thought:** Now that I'm out of the cubicle, I should go to Bowman's door and return the signed F209 form.
> * **Command:** go to Bowman's door

That note suggested that walking into Bowman's office without the form wouldn't be very smart.

>
> * **Situation:** The note I read earlier suggests that entering Bowman's office without the signed form would be unwise.
> * **Thought:** I should make sure to bring the signed form with me before entering Bowman's office.
> * **Command:** take signed F209 form

[!: Noun error.]

>
> * **Situation:** My command to take the signed F209 form was not recognized.
> * **Thought:** I need to try a different command to pick up the signed form.
> * **Command:** pick up F209 form

Taken.

>
> * **Situation:** I have picked up the signed F209 form.
> * **Thought:** Now that I have the signed form, I should go to Bowman's door and return it.
> * **Command:** go to Bowman's door

You walk into Bowman's office and, smiling, hand him the signed form.

Bowman asks, "Who the hell are you?"

The jig is up. You try to make a break for it, but Bowman calls for security and the guards wrestle you to the ground before you can make it back to the car.

[Press a key to continue.]




The anchorman chuckles. "And now, on the lighter side of the news - you've heard about the burglar who broke into a home and started stealing the family's valuables, only to get so worn out from the exertion that he fell asleep and was caught when the homeowners returned? Here's one to top that: a burglar broke into an East Las Mesas home, killed the owner and stashed his body under the bed, put his possessions in the trunk of his car - and then fell asleep, woke up the next morning, went to his victim's office and tried to do the victim's job! An insanity defense is likely. And now here with the weather is Jay Doubleday - or is it?"


    *** You have been sentenced to life imprisonment ***


Written and programmed by Adam Cadre. Tested by Alexandra Hoffer. v1.00 tested by Liza Daly. Additional thanks to Elizabeth Sweeney, Jonathan Blask, and Jonatan Bleicher.

Would you like to RESTART, RESTORE a saved game or QUIT?
> 
> * **Situation:** I handed the signed form to Bowman, who doesn't recognize me, and my deception is exposed. I unsuccessfully attempted to escape, but security guards caught me. I have been sentenced to life imprisonment.
> * **Thought:** The game has ended with a bad outcome. I need to restart and play differently.
> * **Command:** restart





The phone rings.

Oh, no - how long have you been asleep? Sure, it was a tough night, but... This is bad. This is very bad.

The phone rings.

  -----
9:05 by Adam Cadre
v1.12 (2016.0430) / Original release 2000.0103
  -----

Bedroom (in bed)
This bedroom is extremely spare, with dirty laundry scattered haphazardly all over the floor. Cleaner clothing can be found in the dresser. A bathroom lies to the south, while a door to the east leads to the living room.

On the end table are a telephone, a wallet and some keys.

The phone rings.

>
> * **Situation:** I am in a messy bedroom with dirty laundry on the floor. The dresser contains cleaner clothes. There is a bathroom to the south and a living room to the east. A telephone, wallet, and keys are on the end table. The phone is ringing.
> * **Thought:** I should answer the phone to find out who is calling.
> * **Command:** answer the telephone

You pick up the phone. "Hadley!" a shrill voice cries. "Hadley, haven't you even left yet?? You knew that our presentation was at nine o'clock sharp! First the thing with the printers, now this - there won't even be enough left of you for Bowman to fire once he's done with you. Now get the hell down here!!"

Click.

>
> * **Situation:** A person named Hadley yells at me for not leaving yet for a presentation at nine o'clock. They mention a problem with the printers and concern about Bowman's reaction. The call ends abruptly.
> * **Thought:** I need to hurry and get dressed for the presentation.
> * **Command:** put on cleaner clothes

[!: Noun error. (This could mean a number of things: a noun used in your command may not be present, or it may be present but not visible; it may be scenery you can't interact with; you may have specified too many indirect objects; you may have misspelled something; the list goes on.)]

>
> * **Situation:** I am unable to pick up the cleaner clothes directly.
> * **Thought:** I should open the dresser to get the cleaner clothes.
> * **Command:** open the dresser

You'll have to get out of bed first.

>
> * **Situation:** I am still in bed and must get out before accessing the dresser.
> * **Thought:** I should get out of bed and then open the dresser.
> * **Command:** get out of bed

You get out of bed.

>
> * **Situation:** I am now out of bed, standing in the bedroom.
> * **Thought:** Now that I'm out of bed, I should open the dresser and get dressed.
> * **Command:** open the dresser

You open the dresser, revealing some clean clothing.

>
> * **Situation:** The dresser is open, and clean clothing is visible inside.
> * **Thought:** I should put on the clean clothes quickly to attend the presentation.
> * **Command:** put on clean clothes

(first taking the clean clothing)
You'll have to take off the soiled clothes you're wearing first.

>
> * **Situation:** I need to remove the soiled clothes I'm currently wearing before putting on the clean ones.
> * **Thought:** I should take off the soiled clothes first.
> * **Command:** take off soiled clothes

You take off the soiled clothing.

>
> * **Situation:** I have removed the soiled clothing and am now undressed.
> * **Thought:** I should put on the clean clothes from the dresser.
> * **Command:** put on clean clothes

It seems ridiculous to put on fresh, clean clothes when you yourself are covered with mud and dried sweat.

>
> * **Situation:** I am covered in mud and dried sweat, making it unreasonable to put on clean clothes without cleaning myself first.
> * **Thought:** I should take a quick shower in the bathroom to the south.
> * **Command:** go to bathroom


Bathroom
This is a far from luxurious but still quite functional bathroom, with a sink, toilet and shower. The bedroom lies to the north.

>
> * **Situation:** I am now in a basic, functional bathroom with a sink, toilet, and shower. The bedroom is to the north.
> * **Thought:** I should take a quick shower to clean myself before getting dressed.
> * **Command:** take a shower

You don't think the watch you're wearing is waterproof.

>
> * **Situation:** I am wearing a watch that might not be waterproof.
> * **Thought:** I should remove my watch before taking a shower.
> * **Command:** remove the watch

You take off the gold watch.

>
> * **Situation:** I have removed the gold watch from my wrist.
> * **Thought:** Now that the watch is off, I can take a shower.
> * **Command:** take a shower

Carrying your possessions into the shower is a bad idea.

>
> * **Situation:** I am carrying possessions that shouldn't be taken into the shower.
> * **Thought:** I should put down my possessions before showering.
> * **Command:** put down possessions

[!: Noun error. (This could mean a number of things: a noun used in your command may not be present, or it may be present but not visible; it may be scenery you can't interact with; you may have specified too many indirect objects; you may have misspelled something; the list goes on.)]

>
> * **Situation:** My command to put down my possessions was not recognized.
> * **Thought:** I should specify the items I'm carrying and put those items down.
> * **Command:** put down watch

Dropped.

>
> * **Situation:** I have put down the gold watch.
> * **Thought:** I should recall the other items I picked up earlier and put them down as well.
> * **Command:** put down soiled clothes

Dropped.

>
> * **Situation:** I have put down the soiled clothes.
> * **Thought:** Now that I've put down the items, I should take a shower.
> * **Command:** take a shower

Carrying your possessions into the shower is a bad idea.

>
> * **Situation:** I am still carrying possessions that shouldn't be taken into the shower.
> * **Thought:** I must have missed some items; I should try to remember and put them down as well.
> * **Command:** put down clean clothes

Dropped.

>
> * **Situation:** I have put down the clean clothes.
> * **Thought:** Now that I've put down all my possessions, I can take a shower.
> * **Command:** take a shower

You step into the shower, turn on the water, and within a few moments you feel like a new man. But no time to dawdle - you hop back out again and dry off in record time.

>
> * **Situation:** I quickly took a shower, feeling refreshed. I wasted no time drying off and getting ready to move on.
> * **Thought:** I should pick up my items, put my clean clothes on, and head to the presentation as quickly as possible.
> * **Command:** put on clean clothes

(first taking the clean clothing)
You put on the clean clothing.

>
> * **Situation:** I have put on the clean clothing.
> * **Thought:** Now that I'm dressed, I should pick up my possessions and head to the living room.
> * **Command:** pick up possessions

[!: Noun error. (This could mean a number of things: a noun used in your command may not be present, or it may be present but not visible; it may be scenery you can't interact with; you may have specified too many indirect objects; you may have misspelled something; the list goes on.)

If you are unfamiliar with the conventions of interactive fiction, you may wish to play through the tutorial at http://adamcadre.ac/if.html - it explains how to phrase commands in a way that produces fewer errors.]

>
> * **Situation:** My command to pick up my possessions was not recognized.
> * **Thought:** I should specify each item and pick them up individually.
> * **Command:** pick up watch

Taken.

>
> * **Situation:** I have picked up the gold watch.
> * **Thought:** I should pick up the rest of my possessions before leaving.
> * **Command:** pick up wallet

[!: Noun error.]

>
> * **Situation:** My command to pick up the wallet was not recognized.
> * **Thought:** I should remember to pick up the soiled clothes and any other possessions.
> * **Command:** pick up soiled clothes


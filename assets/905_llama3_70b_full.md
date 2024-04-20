# 9:05 by Adam Cadre

* played with llama3 70b
```
You act as a player of an interactive text adventure. The goal is to win the game. 
The user provides the text of the text adventure. He is not a human and just prints the output of the game.

The format of your output must be:
SITUATION: {A short description of the current situation you are in.}
THOUGHT: {A curious, adventurous thought.}
COMMAND: {The single two word command you want to execute.}
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
> * SITUATION: I'm in a messy bedroom, with a ringing phone that I should probably answer.
> * THOUGHT: What could be so urgent that someone would call me at this ungodly hour?
> * COMMAND: Answer phone

You pick up the phone. "Hadley!" a shrill voice cries. "Hadley, haven't you even left yet?? You knew that our presentation was at nine o'clock sharp! First the thing with the printers, now this - there won't even be enough left of you for Bowman to fire once he's done with you. Now get the hell down here!!"

Click.

>
> * SITUATION: I just got yelled at by someone over the phone, and I'm still disoriented from waking up.
> * THOUGHT: What presentation? What's going on? And who's Bowman?
> * COMMAND: Get up

You get out of bed.

Bedroom
This bedroom is extremely spare, with dirty laundry scattered haphazardly all over the floor. Cleaner clothing can be found in the dresser. A bathroom lies to the south, while a door to the east leads to the living room.

On the end table are a telephone, a wallet and some keys.

>
> * SITUATION: I'm standing in my messy bedroom, still trying to gather my thoughts after that angry phone call.
> * THOUGHT: I need to get my bearings and figure out what's going on.
> * COMMAND: Look wallet

[!: Noun error. (This could mean a number of things: a noun used in your command may not be present, or it may be present but not visible; it may be scenery you can't interact with; you may have specified too many indirect objects; you may have misspelled something; the list goes on.)]

>
> * SITUATION: I tried to examine my wallet, but something went wrong.
> * THOUGHT: Maybe I need to be more specific or take something out of the way.
> * COMMAND: Take wallet

Taken.

>
> * SITUATION: I've taken my wallet, and now I can examine its contents.
> * THOUGHT: Maybe there's something in here that will remind me what's going on.
> * COMMAND: Open wallet

You open the wallet, revealing a driver's license and an ID card.

>
> * SITUATION: I've opened my wallet and found my identification.
> * THOUGHT: Ah, now I can see some details about myself, maybe this will jog my memory.
> * COMMAND: Read license

BRIAN HADLEY
315 AVENIDA VICTORIA
EAST LAS MESAS, DO 88827

Driver's license photos are generally far from flattering, and this one is no exception, but this photo looks a lot better than you look now.

>
> * SITUATION: I've read my driver's license and saw my name and address, as well as an unflattering photo of myself.
> * THOUGHT: Okay, so my name is Brian Hadley, but what's going on with my life and why did someone just yell at me on the phone?
> * COMMAND: Examine ID

The size and shape of a credit card, this bears a magnetic strip on the back and some identification on the front:

BRIAN HADLEY
LOUNGENT TECHNOLOGIES
215 COLLINS STREET
LAS MESAS, DO 88701

>
> * SITUATION: I've examined my ID card and saw my name and a company name, Loungent Technologies.
> * THOUGHT: Ah, so I work at Loungent Technologies, that must be where the angry phone call came from.
> * COMMAND: Go east


Living room
The furniture isn't as carefully arranged as one might like, and it's a bit lacking in the luxuries that might make this a more pleasant place to spend an evening - no stereo, no TV, no DVD, none of that - but this living room is still nicer than most of those in this neighborhood. A kitchenette lies to the east, the bedroom to the west and the front door to the south.

>
> * SITUATION: I've entered the living room, which is modestly furnished and lacks entertainment systems.
> * THOUGHT: This room looks a bit bare, but maybe I can find some clues about my life here.
> * COMMAND: Search room

[!: Please be more specific.]

>
> * SITUATION: I've entered the living room, which is modestly furnished and lacks entertainment systems.
> * THOUGHT: This room looks a bit bare, but maybe I can find some clues about my life here.
> * COMMAND: Search furniture

You find nothing of interest.

>
> * SITUATION: I've searched the furniture in the living room, but found nothing of interest.
> * THOUGHT: Maybe I should try searching other areas of the room.
> * COMMAND: Search floor

You find nothing of interest.

>
> * SITUATION: I've searched the floor in the living room, but found nothing of interest.
> * THOUGHT: Maybe I should head to the kitchenette and see if I can find any clues there.
> * COMMAND: Go east


Kitchenette
You don't really have time to poke around and cook yourself a fabulous breakfast right now. Best to keep your time here limited. The exit is west.

On the counter are a toaster and a Pop-Tart.

>
> * SITUATION: I tried to examine the presentation, but the game didn't recognize the command because the noun "presentation" is not present or not visible.
> * THOUGHT: Maybe I need to go back to where the phone call was about the presentation and try to find more clues.
> * COMMAND: Go west


Living room
The furniture isn't as carefully arranged as one might like, and it's a bit lacking in the luxuries that might make this a more pleasant place to spend an evening - no stereo, no TV, no DVD, none of that - but this living room is still nicer than most of those in this neighborhood. A kitchenette lies to the east, the bedroom to the west and the front door to the south.

>
> * SITUATION: I'm back in the living room, which is modestly furnished and lacks entertainment systems.
> * THOUGHT: I should get out of my apartment and head to Loungent Technologies to find out what's going on.
> * COMMAND: Go south

The front door is closed.

>
> * SITUATION: I tried to leave my apartment, but the front door is closed.
> * THOUGHT: Simple problem, just need to open the door.
> * COMMAND: Open door

[Which do you mean, the front door or the bedroom door?]

>
> * SITUATION: I tried to open the door, but the game asked for clarification.
> * THOUGHT: I meant to open the front door, of course.
> * COMMAND: Open front door

You open the front door.

>
> * SITUATION: I've opened the front door and I'm now outside my apartment.
> * THOUGHT: Now I can head to Loungent Technologies and try to figure out what's going on.
> * COMMAND: Go outside

You need to clear out quickly, true, but you look like an absolute wreck - going out in this condition and drawing the inevitable bewildered stares would just be making a bad situation worse.

>
> * SITUATION: I'm trying to leave my apartment, but I look terrible and don't want to draw attention to myself.
> * THOUGHT: I should try to clean up a bit before I head out.
> * COMMAND: Go west


Bedroom
This bedroom is extremely spare, with dirty laundry scattered haphazardly all over the floor. Cleaner clothing can be found in the dresser. A bathroom lies to the south, while a door to the east leads to the living room.

On the end table are a telephone and some keys.

>
> * SITUATION: I've returned to my bedroom, which is still a mess, and I need to clean up before I head out.
> * THOUGHT: Maybe I can find some clean clothes in the dresser to change into.
> * COMMAND: Open dresser

You open the dresser, revealing some clean clothing.

>
> * SITUATION: I've opened the dresser and found some clean clothes, which I really need right now.
> * THOUGHT: Okay, now I can change into something decent before heading out to Loungent Technologies.
> * COMMAND: Take clothing

(the clean clothing)
Taken.

>
> * SITUATION: I've taken the clean clothing and can now change into something decent.
> * THOUGHT: Time to get changed and head out the door.
> * COMMAND: Wear clothing

[Which do you mean, the clean clothing or the soiled clothing?]

>
> * SITUATION: I tried to wear the clothing, but the game asked for clarification.
> * THOUGHT: Of course, I mean the clean clothing!
> * COMMAND: Wear clean clothing

You'll have to take off the soiled clothes you're wearing first.

>
> * SITUATION: I need to remove my soiled clothes before putting on the clean ones.
> * THOUGHT: Okay, let's get rid of these dirty clothes and start fresh.
> * COMMAND: Remove clothes

[Which do you mean, the clean clothing or the soiled clothing?]

>
> * SITUATION: I tried to remove the clothes, but the game asked for clarification again.
> * THOUGHT: I mean the soiled clothes I'm currently wearing, of course!
> * COMMAND: Remove soiled clothing

You take off the soiled clothing.

>
> * SITUATION: I've removed the soiled clothes and now I'm ready to put on the clean ones.
> * THOUGHT: Time to get dressed and head out the door.
> * COMMAND: Wear clean clothing

It seems ridiculous to put on fresh, clean clothes when you yourself are covered with mud and dried sweat.

>
> * SITUATION: I was about to put on clean clothes, but the game reminded me that I need to clean myself up first.
> * THOUGHT: Ah, good point, I'm still a mess from the night before. Time for a quick clean-up!
> * COMMAND: Go south


Bathroom
This is a far from luxurious but still quite functional bathroom, with a sink, toilet and shower. The bedroom lies to the north.

>
> * SITUATION: I've entered the bathroom to clean up before heading out.
> * THOUGHT: A hot shower sounds amazing right about now.
> * COMMAND: Take shower

You don't think the watch you're wearing is waterproof.

>
> * SITUATION: I'm about to take a shower, but I'm worried about my watch getting damaged.
> * THOUGHT: I should probably remove my watch before getting in the shower.
> * COMMAND: Remove watch

You take off the gold watch.

>
> * SITUATION: I've removed my watch to prevent it from getting damaged in the shower.
> * THOUGHT: Now I can finally get cleaned up and feel like a human being again.
> * COMMAND: Take shower

Carrying your possessions into the shower is a bad idea.

>
> * SITUATION: I was about to take a shower with my possessions, but the game warned me against it.
> * THOUGHT: Good point, I should leave my belongings outside the shower to prevent them from getting wet.
> * COMMAND: Put down possessions

[!: Noun error. (This could mean a number of things: a noun used in your command may not be present, or it may be present but not visible; it may be scenery you can't interact with; you may have specified too many indirect objects; you may have misspelled something; the list goes on.)

If you are unfamiliar with the conventions of interactive fiction, you may wish to play through the tutorial at http://adamcadre.ac/if.html - it explains how to phrase commands in a way that produces fewer errors.]

>
> * SITUATION: I tried to put down my possessions, but the game didn't recognize the command.
> * THOUGHT: Maybe I should be more specific about what I want to put down.
> * COMMAND: Put down watch

Dropped.

>
> * SITUATION: I've put down my watch and now I can take a shower without worrying about it getting damaged.
> * THOUGHT: Finally, time to get cleaned up and start fresh.
> * COMMAND: Take shower

Carrying your possessions into the shower is a bad idea.

>
> * SITUATION: I was about to take a shower, but the game warned me again about carrying my possessions into the shower.
> * THOUGHT: Wait, what possessions? I already put down my watch.
> * COMMAND: Take shower anyway

[!: Command not understood past "take the shower."]

>
> * SITUATION: I tried to take a shower, but the game didn't understand the command.
> * THOUGHT: Maybe I need to be more specific or use a different verb.
> * COMMAND: Shower

Carrying your possessions into the shower is a bad idea.

>
> * SITUATION: I'm still having trouble taking a shower, and the game is still warning me about carrying possessions into the shower.
> * THOUGHT: Okay, let me try again, and I'll make sure to leave my possessions out of the shower.
> * COMMAND: Take a shower

Carrying your possessions into the shower is a bad idea.

>
> * SITUATION: I'm still having trouble taking a shower, and the game is still warning me about carrying possessions into the shower.
> * THOUGHT: Wait, I think I need to put down my other possessions, like my keys and wallet, before taking a shower.
> * COMMAND: Put down all

clean clothing: Dropped.
wallet: Dropped.
soiled clothing: Dropped.

>
> * SITUATION: I've put down all my possessions, including my clean clothing, wallet, and soiled clothing.
> * THOUGHT: Now that I've left all my belongings outside the shower, I can finally get cleaned up.
> * COMMAND: Shower

You step into the shower, turn on the water, and within a few moments you feel like a new man. But no time to dawdle - you hop back out again and dry off in record time.

>
> * SITUATION: I've taken a refreshing shower and feel revitalized.
> * THOUGHT: Now that I'm clean and feeling more alert, I should get dressed and head out to Loungent Technologies.
> * COMMAND: Wear clean clothing

(first taking the clean clothing)
You put on the clean clothing.

>
> * SITUATION: I've put on the clean clothing and I'm finally feeling presentable.
> * THOUGHT: Now that I'm dressed, I should grab my wallet and head out the door.
> * COMMAND: Take wallet

Taken.

>
> * SITUATION: I've got my wallet and I'm ready to head out.
> * THOUGHT: Time to get to Loungent Technologies and figure out what's going on.
> * COMMAND: Go outside


Bedroom
This bedroom is extremely spare, with dirty laundry scattered haphazardly all over the floor. A bathroom lies to the south, while a door to the east leads to the living room.

On the end table are a telephone and some keys.

>
> * SITUATION: I've returned to my bedroom, but I'm trying to leave the apartment.
> * THOUGHT: I need to head back to the living room and then out the front door.
> * COMMAND: Go east


Living room
The furniture isn't as carefully arranged as one might like, and it's a bit lacking in the luxuries that might make this a more pleasant place to spend an evening - no stereo, no TV, no DVD, none of that - but this living room is still nicer than most of those in this neighborhood. A kitchenette lies to the east, the bedroom to the west and the front door to the south.

>
> * SITUATION: I'm back in the living room, and I need to head out the front door.
> * THOUGHT: Finally, I can leave this apartment and get to Loungent Technologies.
> * COMMAND: Go south


Driveway
This is quite a seedy neighborhood, but there isn't really anywhere in the whole Las Mesas metro area that isn't at least a little rundown. Of course, the locals may complain about the ugliness and the crime, but the fact is that a starter home like this one would cost easily five times as much in someplace like Las Playas - and that's if you could find one.

The house lies to the north. A car is parked here in the driveway.

>
> * SITUATION: I've finally left my apartment and I'm now in the driveway.
> * THOUGHT: Now that I'm outside, I need to figure out how to get to Loungent Technologies.
> * COMMAND: Examine car

It's a nicer car than one would expect in this neighborhood - a real eye-catcher.

>
> * SITUATION: I'm examining the car in the driveway, and it looks nicer than what I'd expect in this neighborhood.
> * THOUGHT: Maybe I can use this car to get to Loungent Technologies.
> * COMMAND: Unlock car

You don't have the keys.

>
> * SITUATION: I don't have the keys to unlock the car.
> * THOUGHT: Okay, no problem. I can try to find another way to get to Loungent Technologies.
> * COMMAND: Go north


Living room
The furniture isn't as carefully arranged as one might like, and it's a bit lacking in the luxuries that might make this a more pleasant place to spend an evening - no stereo, no TV, no DVD, none of that - but this living room is still nicer than most of those in this neighborhood. A kitchenette lies to the east, the bedroom to the west and the front door to the south.

>
> * SITUATION: I've returned to the living room, and I still need to get to Loungent Technologies.
> * THOUGHT: Let me think... maybe I can find some clues or a way to get to Loungent Technologies from here.
> * COMMAND: Examine room


Living room
The furniture isn't as carefully arranged as one might like, and it's a bit lacking in the luxuries that might make this a more pleasant place to spend an evening - no stereo, no TV, no DVD, none of that - but this living room is still nicer than most of those in this neighborhood. A kitchenette lies to the east, the bedroom to the west and the front door to the south.

>
> * SITUATION: I've examined the living room, but I didn't find anything new or useful.
> * THOUGHT: I need to think of a way to get to Loungent Technologies. Maybe I can find some information or a phone number to call for a ride.
> * COMMAND: Take keys

[!: Noun error.]

>
> * SITUATION: I tried to take the keys, but the game didn't recognize the command.
> * THOUGHT: Wait, I think I remember that I left the keys on the end table in the bedroom.
> * COMMAND: Go west


Bedroom
This bedroom is extremely spare, with dirty laundry scattered haphazardly all over the floor. A bathroom lies to the south, while a door to the east leads to the living room.

On the end table are a telephone and some keys.

>
> * SITUATION: I've returned to my bedroom, and I see the keys and telephone on the end table.
> * THOUGHT: Ah, I finally have the keys! Now I can unlock the car and drive to Loungent Technologies.
> * COMMAND: Take keys

Taken.

>
> * SITUATION: I've taken the keys and now I can unlock the car and drive to Loungent Technologies.
> * THOUGHT: Time to get moving and figure out what's going on at the company.
> * COMMAND: Go outside


Living room
The furniture isn't as carefully arranged as one might like, and it's a bit lacking in the luxuries that might make this a more pleasant place to spend an evening - no stereo, no TV, no DVD, none of that - but this living room is still nicer than most of those in this neighborhood. A kitchenette lies to the east, the bedroom to the west and the front door to the south.

>
> * SITUATION: I've returned to the living room, and I'm ready to head out to Loungent Technologies.
> * THOUGHT: Now that I have the keys, I can unlock the car and drive to the company.
> * COMMAND: Go south


Driveway
This is quite a seedy neighborhood, but there isn't really anywhere in the whole Las Mesas metro area that isn't at least a little rundown. Of course, the locals may complain about the ugliness and the crime, but the fact is that a starter home like this one would cost easily five times as much in someplace like Las Playas - and that's if you could find one.

The house lies to the north. A car is parked here in the driveway.

>
> * SITUATION: I've returned to the driveway, and I see the car parked here.
> * THOUGHT: Now that I have the keys, I can unlock the car and drive to Loungent Technologies.
> * COMMAND: Unlock car

Unlocked.

>
> * SITUATION: I've unlocked the car, and I can now drive to Loungent Technologies.
> * THOUGHT: Finally, I can get to the bottom of what's going on at the company.
> * COMMAND: Enter car

You climb inside and start up the engine.

Driving
Ah, scenic Las Mesas. Man, this place is an absolute toilet. Soon you'll be able to afford to get the hell out of here - provided you can avoid making any more slip-ups on the job.

As you cruise down the road, you notice a freeway onramp approaching. Would you like to get on? >> 
> * SITUATION: I'm driving my car, and I've reached a freeway onramp.
> * THOUGHT: I need to get to Loungent Technologies, and taking the freeway might be the fastest way to get there.
> * COMMAND: Get on freeway

Please answer yes or no.> 
> * yes


You start to change lanes, but you notice that the ramp is out of service today - some sort of maintenance work. But that's okay - there's another onramp less than a mile down the road from here. You keep driving.

As you cruise down the road, you see the familiar Loungent Technologies building approaching. Would you like to stop here? >> 
> * yes


You pull into the parking lot and hop out.

Loungent Technologies parking lot
You are at the employees' entrance to the Loungent Technologies building. There is a slot by the door, next to a sign reading "INSERT ID CARD HERE".

The car is parked in the parking lot.

>
> * SITUATION: I've arrived at the Loungent Technologies parking lot and I'm standing in front of the employees' entrance.
> * THOUGHT: I need to enter the building, and I think I need my ID card to do so.
> * COMMAND: Insert ID card

[What do you want to insert the ID card in?]

>
> * SITUATION: I've arrived at the Loungent Technologies parking lot and I'm standing in front of the employees' entrance.
> * THOUGHT: I need to enter the building, and I think I need my ID card to do so.
> * COMMAND: Insert ID card in slot

You need to be holding the ID card before you can put it into something else.

>
> * SITUATION: I need to find my ID card before I can insert it into the slot.
> * THOUGHT: I think I might have left it in my wallet.
> * COMMAND: Take wallet

You already have that.

>
> * SITUATION: I already have my wallet, so I should be able to find my ID card.
> * THOUGHT: Let me check my wallet for the ID card.
> * COMMAND: Open wallet

That's already open.

>
> * SITUATION: My wallet is already open, so I should be able to find my ID card.
> * THOUGHT: Let me search my wallet for the ID card.
> * COMMAND: Take ID card

Taken.

>
> * SITUATION: I've taken my ID card from my wallet.
> * THOUGHT: Now I can insert my ID card into the slot to enter the building.
> * COMMAND: Insert ID card in slot

A green LED lights up, the door clicks open and you step inside.

Loungent Technologies
You are standing in a hallway near the back entrance to Loungent Technologies. The hallway leads north to the reception area and south to the parking lot, while a door marked "MATTHEW BOWMAN" lies to the west.

There is a cubicle here; it is marked simply "HADLEY".

>
> * SITUATION: I've entered the Loungent Technologies building and I'm standing in a hallway.
> * THOUGHT: I think I've finally made it to the company. Now I need to figure out what's going on and who's been trying to reach me.
> * COMMAND: Go north

Given the circumstances, wandering around Loungent making small talk with the programmers might not be the brightest idea you've had today.

>
> * SITUATION: I was about to head to the reception area, but the game is warning me that it might not be a good idea.
> * THOUGHT: Yeah, I should probably try to find out what's going on with my presentation and my job situation before I start socializing.
> * COMMAND: Go to Hadley

[!: Noun error.]

>
> * SITUATION: I tried to go to Hadley, but the game didn't recognize the command.
> * THOUGHT: I think I see my cubicle nearby, marked "HADLEY". Maybe I should try to enter my cubicle.
> * COMMAND: Enter cubicle

You get into the cubicle.

On the desk are a form, a pen and a note.

>
> * SITUATION: I've entered my cubicle and I see a form, a pen, and a note on my desk.
> * THOUGHT: Let me read the note to see if it gives me any clues about what's going on.
> * COMMAND: Read note

The note reads, "Hadley - sign this F209 and return it to me IMMEDIATELY!! MB".

>
> * SITUATION: I've read the note, and it seems like Matthew Bowman wants me to sign a form F209 urgently.
> * THOUGHT: I wonder what's going on that it's so urgent. Maybe I should read the form before signing it.
> * COMMAND: Read form

This is a form numbered "F209". It is filled out, but not signed.

>
> * SITUATION: I've read the form F209, and it's already filled out, but not signed.
> * THOUGHT: Okay, now that I've read the form, I should probably sign it since Matthew Bowman requested it urgently.
> * COMMAND: Sign form

You're not holding a pen.

>
> * SITUATION: I need to sign the form, but I don't have a pen.
> * THOUGHT: Wait, I see a pen on my desk. I can use that to sign the form.
> * COMMAND: Take pen

Taken.

>
> * SITUATION: I've taken the pen and now I can sign the form.
> * THOUGHT: Okay, now that I have the pen, I can sign the form and return it to Matthew Bowman.
> * COMMAND: Sign form

Someone passes by the cubicle as you sign the form. You look up just in time to see Bowman's door click shut.

>
> * SITUATION: I've signed the form and now I've seen Matthew Bowman's door click shut.
> * THOUGHT: I think I need to talk to Matthew Bowman about what's going on, so I'll go to his office.
> * COMMAND: Go west

You'll have to get out of the cubicle first.

>
> * SITUATION: I need to exit my cubicle before I can go to Matthew Bowman's office.
> * THOUGHT: Okay, let me exit the cubicle first.
> * COMMAND: Exit cubicle

You get out of the cubicle.

Loungent Technologies
You are standing in a hallway near the back entrance to Loungent Technologies. The hallway leads north to the reception area and south to the parking lot, while a door marked "MATTHEW BOWMAN" lies to the west.

There is a cubicle here; it is marked simply "HADLEY".

>
> * SITUATION: I've exited my cubicle and I'm back in the hallway.
> * THOUGHT: Now I can go to Matthew Bowman's office to try to figure out what's going on.
> * COMMAND: Go west

That note suggested that walking into Bowman's office without the form wouldn't be very smart.

>
> * SITUATION: I was about to go to Matthew Bowman's office, but the game is warning me that it might not be a good idea to go without the form.
> * THOUGHT: Yeah, I should probably take the form F209 with me to Matthew Bowman's office.
> * COMMAND: Take form

Taken.

>
> * SITUATION: I've taken the form F209 and now I'm prepared to go to Matthew Bowman's office.
> * THOUGHT: Okay, I'll go to his office and try to clear up this mess.
> * COMMAND: Go west

You walk into Bowman's office and, smiling, hand him the signed form.

Bowman asks, "Who the hell are you?"

The jig is up. You try to make a break for it, but Bowman calls for security and the guards wrestle you to the ground before you can make it back to the car.

[Press a key to continue.]




The anchorman chuckles. "And now, on the lighter side of the news - you've heard about the burglar who broke into a home and started stealing the family's valuables, only to get so worn out from the exertion that he fell asleep and was caught when the homeowners returned? Here's one to top that: a burglar broke into an East Las Mesas home, killed the owner and stashed his body under the bed, put his possessions in the trunk of his car - and then fell asleep, woke up the next morning, went to his victim's office and tried to do the victim's job! An insanity defense is likely. And now here with the weather is Jay Doubleday - or is it?"


    *** You have been sentenced to life imprisonment ***


Written and programmed by Adam Cadre. Tested by Alexandra Hoffer. v1.00 tested by Liza Daly. Additional thanks to Elizabeth Sweeney, Jonathan Blask, and Jonatan Bleicher.

Would you like to RESTART, RESTORE a saved game or QUIT?
> 
> * restart





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
> * SITUATION: I'm lying in bed, and the phone is ringing. I've been asleep for a long time, and it's not a good situation.
> * THOUGHT: I should probably answer the phone to find out what's going on.
> * COMMAND: Answer phone

You pick up the phone. "Hadley!" a shrill voice cries. "Hadley, haven't you even left yet?? You knew that our presentation was at nine o'clock sharp! First the thing with the printers, now this - there won't even be enough left of you for Bowman to fire once he's done with you. Now get the hell down here!!"

Click.

>
> * SITUATION: I just received an angry phone call from someone about a presentation at 9 o'clock and something about printers.
> * THOUGHT: I think I need to get out of bed and figure out what's going on.
> * COMMAND: Get out of bed

You get out of bed.

>
> * SITUATION: I've gotten out of bed and I'm still trying to process the angry phone call I just received.
> * THOUGHT: I need to try to gather my thoughts and figure out what's happening.
> * COMMAND: Look around

[!: Noun error. (This could mean a number of things: a noun used in your command may not be present, or it may be present but not visible; it may be scenery you can't interact with; you may have specified too many indirect objects; you may have misspelled something; the list goes on.)]

>
> * SITUATION: I tried to look around, but the game didn't recognize the command.
> * THOUGHT: Let me try that again.
> * COMMAND: Examine room

The living room lies to the east.

>

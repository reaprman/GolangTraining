package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bs)
}

/*
Run Result:
"The aorta of a whale is larger in the bore than the main pipe of the
water-works at London Bridge, and the water roaring in its passage
through that pipe is inferior in impetus and velocity to the blood
gushing from the whale's heart." --PALEY'S THEOLOGY.

"The whale is a mammiferous animal without hind feet." --BARON
CUVIER.

"In 40 degrees south, we saw Spermacetti Whales, but did not take any
till the first of May, the sea being then covered with them."
--COLNETT'S VOYAGE FOR THE PURPOSE OF EXTENDING THE SPERMACETI WHALE
FISHERY.

"In the free element beneath me swam,
Floundered and dived, in play, in chace, in battle,
Fishes of every colour, form, and kind;
Which language cannot paint, and mariner
Had never seen; from dread Leviathan
To insect millions peopling every wave:
Gather'd in shoals immense, like floating islands,
Led by mysterious instincts through that waste
And trackless region, though on every side
Assaulted by voracious enemies,
Whales, sharks, and monsters, arm'd in front or jaw,
With swords, saws, spiral horns, or hooked fangs."
--MONTGOMERY'S WORLD BEFORE THE FLOOD.

"Io!  Paean!  Io! sing.
To the finny people's king.
Not a mightier whale than this
In the vast Atlantic is;
Not a fatter fish than he,
Flounders round the Polar Sea." --CHARLES LAMB'S TRIUMPH OF THE
WHALE.

"In the year 1690 some persons were on a high hill observing the
whales spouting and sporting with each other, when one observed:
there--pointing to the sea--is a green pasture where our children's
grand-children will go for bread." --OBED MACY'S HISTORY OF
NANTUCKET.

"I built a cottage for Susan and myself and made a gateway in the
form of a Gothic Arch, by setting up a whale's jaw bones."
--HAWTHORNE'S TWICE TOLD TALES.

"She came to bespeak a monument for her first love, who had been
killed by a whale in the Pacific ocean, no less than forty years
ago." --IBID.

"No, Sir, 'tis a Right Whale," answered Tom; "I saw his sprout; he
threw up a pair of as pretty rainbows as a Christian would wish to
look at.  He's a raal oil-butt, that fellow!" --COOPER'S PILOT.

"The papers were brought in, and we saw in the Berlin Gazette that
whales had been introduced on the stage there." --ECKERMANN'S
CONVERSATIONS WITH GOETHE.

"My God!  Mr. Chace, what is the matter?"  I answered, "we have been
stove by a whale." --"NARRATIVE OF THE SHIPWRECK OF THE WHALE SHIP
ESSEX OF NANTUCKET, WHICH WAS ATTACKED AND FINALLY DESTROYED BY A
LARGE SPERM WHALE IN THE PACIFIC OCEAN."  BY OWEN CHACE OF NANTUCKET,
FIRST MATE OF SAID VESSEL.  NEW YORK, 1821.

"A mariner sat in the shrouds one night,
The wind was piping free;
Now bright, now dimmed, was the moonlight pale,
And the phospher gleamed in the wake of the whale,
As it floundered in the sea." --ELIZABETH OAKES SMITH.

"The quantity of line withdrawn from the boats engaged in the capture
of this one whale, amounted altogether to 10,440 yards or nearly six
English miles. ...

"Sometimes the whale shakes its tremendous tail in the air, which,
cracking like a whip, resounds to the distance of three or four
miles." --SCORESBY.

"Mad with the agonies he endures from these fresh attacks, the
infuriated Sperm Whale rolls over and over; he rears his enormous
head, and with wide expanded jaws snaps at everything around him; he
rushes at the boats with his head; they are propelled before him with
vast swiftness, and sometimes utterly destroyed. ...  It is a matter
of great astonishment that the consideration of the habits of so
interesting, and, in a commercial point of view, so important an
animal (as the Sperm Whale) should have been so entirely neglected,
or should have excited so little curiosity among the numerous, and
many of them competent observers, that of late years, must have
possessed the most abundant and the most convenient opportunities of
witnessing their habitudes." --THOMAS BEALE'S HISTORY OF THE SPERM
WHALE, 1839.

"The Cachalot" (Sperm Whale) "is not only better armed than the True
Whale" (Greenland or Right Whale) "in possessing a formidable weapon
at either extremity of its body, but also more frequently displays a
disposition to employ these weapons offensively and in manner at once
so artful, bold, and mischievous, as to lead to its being regarded as
the most dangerous to attack of all the known species of the whale
tribe." --FREDERICK DEBELL BENNETT'S WHALING VOYAGE ROUND THE GLOBE,
1840.

October 13.  "There she blows," was sung out from the mast-head.
"Where away?" demanded the captain.
"Three points off the lee bow, sir."
"Raise up your wheel.  Steady!"  "Steady, sir."
"Mast-head ahoy!  Do you see that whale now?"
"Ay ay, sir!  A shoal of Sperm Whales!  There she blows!  There she
breaches!"
"Sing out! sing out every time!"
"Ay Ay, sir!  There she blows! there--there--THAR she
blows--bowes--bo-o-os!"
"How far off?"
"Two miles and a half."
"Thunder and lightning! so near!  Call all hands." --J. ROSS BROWNE'S
ETCHINGS OF A WHALING CRUIZE.  1846.

"The Whale-ship Globe, on board of which vessel occurred the horrid
transactions we are about to relate, belonged to the island of
Nantucket." --"NARRATIVE OF THE GLOBE," BY LAY AND HUSSEY SURVIVORS.
A.D. 1828.

Being once pursued by a whale which he had wounded, he parried the
assault for some time with a lance; but the furious monster at length
rushed on the boat; himself and comrades only being preserved by
leaping into the water when they saw the onset was inevitable."
--MISSIONARY JOURNAL OF TYERMAN AND BENNETT.

"Nantucket itself," said Mr. Webster, "is a very striking and
peculiar portion of the National interest.  There is a population of
eight or nine thousand persons living here in the sea, adding largely
every year to the National wealth by the boldest and most persevering
industry." --REPORT OF DANIEL WEBSTER'S SPEECH IN THE U.  S.  SENATE,
ON THE APPLICATION FOR THE ERECTION OF A BREAKWATER AT NANTUCKET.
1828.

"The whale fell directly over him, and probably killed him in a
moment." --"THE WHALE AND HIS CAPTORS, OR THE WHALEMAN'S ADVENTURES
AND THE WHALE'S BIOGRAPHY, GATHERED ON THE HOMEWARD CRUISE OF THE
COMMODORE PREBLE."  BY REV. HENRY T. CHEEVER.

"If you make the least damn bit of noise," replied Samuel, "I will
send you to hell." --LIFE OF SAMUEL COMSTOCK (THE MUTINEER), BY HIS
BROTHER, WILLIAM COMSTOCK.  ANOTHER VERSION OF THE WHALE-SHIP GLOBE
NARRATIVE.

"The voyages of the Dutch and English to the Northern Ocean, in
order, if possible, to discover a passage through it to India, though
they failed of their main object, laid-open the haunts of the whale."
--MCCULLOCH'S COMMERCIAL DICTIONARY.

"These things are reciprocal; the ball rebounds, only to bound
forward again; for now in laying open the haunts of the whale, the
whalemen seem to have indirectly hit upon new clews to that same
mystic North-West Passage." --FROM "SOMETHING" UNPUBLISHED.

"It is impossible to meet a whale-ship on the ocean without being
struck by her near appearance.  The vessel under short sail, with
look-outs at the mast-heads, eagerly scanning the wide expanse around
them, has a totally different air from those engaged in regular
voyage." --CURRENTS AND WHALING.  U.S. EX. EX.

"Pedestrians in the vicinity of London and elsewhere may recollect
having seen large curved bones set upright in the earth, either to
form arches over gateways, or entrances to alcoves, and they may
perhaps have been told that these were the ribs of whales." --TALES
OF A WHALE VOYAGER TO THE ARCTIC OCEAN.

"It was not till the boats returned from the pursuit of these whales,
that the whites saw their ship in bloody possession of the savages
enrolled among the crew." --NEWSPAPER ACCOUNT OF THE TAKING AND
RETAKING OF THE WHALE-SHIP HOBOMACK.

"It is generally well known that out of the crews of Whaling vessels
(American) few ever return in the ships on board of which they
departed." --CRUISE IN A WHALE BOAT.

"Suddenly a mighty mass emerged from the water, and shot up
perpendicularly into the air.  It was the while." --MIRIAM COFFIN OR
THE WHALE FISHERMAN.

"The Whale is harpooned to be sure; but bethink you, how you would
manage a powerful unbroken colt, with the mere appliance of a rope
tied to the root of his tail." --A CHAPTER ON WHALING IN RIBS AND
TRUCKS.

"On one occasion I saw two of these monsters (whales) probably male
and female, slowly swimming, one after the other, within less than a
stone's throw of the shore" (Terra Del Fuego), "over which the beech
tree extended its branches." --DARWIN'S VOYAGE OF A NATURALIST.

"'Stern all!' exclaimed the mate, as upon turning his head, he saw
the distended jaws of a large Sperm Whale close to the head of the
boat, threatening it with instant destruction;--'Stern all, for your
lives!'" --WHARTON THE WHALE KILLER.

"So be cheery, my lads, let your hearts never fail,
While the bold harpooneer is striking the whale!" --NANTUCKET SONG.

"Oh, the rare old Whale, mid storm and gale
In his ocean home will be
A giant in might, where might is right,
And King of the boundless sea." --WHALE SONG.



CHAPTER 1

Loomings.


Call me Ishmael.  Some years ago--never mind how long
precisely--having little or no money in my purse, and nothing
particular to interest me on shore, I thought I would sail about a
little and see the watery part of the world.  It is a way I have of
driving off the spleen and regulating the circulation.  Whenever I
find myself growing grim about the mouth; whenever it is a damp,
drizzly November in my soul; whenever I find myself involuntarily
pausing before coffin warehouses, and bringing up the rear of every
funeral I meet; and especially whenever my hypos get such an upper
hand of me, that it requires a strong moral principle to prevent me
from deliberately stepping into the street, and methodically knocking
people's hats off--then, I account it high time to get to sea as soon
as I can.  This is my substitute for pistol and ball.  With a
philosophical flourish Cato throws himself upon his sword; I quietly
take to the ship.  There is nothing surprising in this.  If they but
knew it, almost all men in their degree, some time or other, cherish
very nearly the same feelings towards the ocean with me.

There now is your insular city of the Manhattoes, belted round by
wharves as Indian isles by coral reefs--commerce surrounds it with
her surf.  Right and left, the streets take you waterward.  Its
extreme downtown is the battery, where that noble mole is washed by
waves, and cooled by breezes, which a few hours previous were out of
sight of land.  Look at the crowds of water-gazers there.

Circumambulate the city of a dreamy Sabbath afternoon.  Go from
Corlears Hook to Coenties Slip, and from thence, by Whitehall,
northward.  What do you see?--Posted like silent sentinels all around
the town, stand thousands upon thousands of mortal men fixed in ocean
reveries.  Some leaning against the spiles; some seated upon the
pier-heads; some looking over the bulwarks of ships from China; some
high aloft in the rigging, as if striving to get a still better
seaward peep.  But these are all landsmen; of week days pent up in
lath and plaster--tied to counters, nailed to benches, clinched to
desks.  How then is this?  Are the green fields gone?  What do they
here?

But look! here come more crowds, pacing straight for the water, and
seemingly bound for a dive.  Strange!  Nothing will content them but
the extremest limit of the land; loitering under the shady lee of
yonder warehouses will not suffice.  No.  They must get just as nigh
the water as they possibly can without falling in.  And there they
stand--miles of them--leagues.  Inlanders all, they come from lanes
and alleys, streets and avenues--north, east, south, and west.  Yet
here they all unite.  Tell me, does the magnetic virtue of the
needles of the compasses of all those ships attract them thither?

Once more.  Say you are in the country; in some high land of lakes.
Take almost any path you please, and ten to one it carries you down
in a dale, and leaves you there by a pool in the stream.  There is
magic in it.  Let the most absent-minded of men be plunged in his
deepest reveries--stand that man on his legs, set his feet a-going,
and he will infallibly lead you to water, if water there be in all
that region.  Should you ever be athirst in the great American
desert, try this experiment, if your caravan happen to be supplied
with a metaphysical professor.  Yes, as every one knows, meditation
and water are wedded for ever.

But here is an artist.  He desires to paint you the dreamiest,
shadiest, quietest, most enchanting bit of romantic landscape in all
the valley of the Saco.  What is the chief element he employs?  There
stand his trees, each with a hollow trunk, as if a hermit and a
crucifix were within; and here sleeps his meadow, and there sleep his
cattle; and up from yonder cottage goes a sleepy smoke.  Deep into
distant woodlands winds a mazy way, reaching to overlapping spurs of
mountains bathed in their hill-side blue.  But though the picture
lies thus tranced, and though this pine-tree shakes down its sighs
like leaves upon this shepherd's head, yet all were vain, unless the
shepherd's eye were fixed upon the magic stream before him.  Go visit
the Prairies in June, when for scores on scores of miles you wade
knee-deep among Tiger-lilies--what is the one charm
wanting?--Water--there is not a drop of water there!  Were Niagara
but a cataract of sand, would you travel your thousand miles to see
it?  Why did the poor poet of Tennessee, upon suddenly receiving two
handfuls of silver, deliberate whether to buy him a coat, which he
sadly needed, or invest his money in a pedestrian trip to Rockaway
Beach?  Why is almost every robust healthy boy with a robust healthy
soul in him, at some time or other crazy to go to sea?  Why upon your
first voyage as a passenger, did you yourself feel such a mystical
vibration, when first told that you and your ship were now out of
sight of land?  Why did the old Persians hold the sea holy?  Why did
the Greeks give it a separate deity, and own brother of Jove?  Surely
all this is not without meaning.  And still deeper the meaning of
that story of Narcissus, who because he could not grasp the
tormenting, mild image he saw in the fountain, plunged into it and
was drowned.  But that same image, we ourselves see in all rivers and
oceans.  It is the image of the ungraspable phantom of life; and this
is the key to it all.

Now, when I say that I am in the habit of going to sea whenever I
begin to grow hazy about the eyes, and begin to be over conscious of
my lungs, I do not mean to have it inferred that I ever go to sea as
a passenger.  For to go as a passenger you must needs have a purse,
and a purse is but a rag unless you have something in it.  Besides,
passengers get sea-sick--grow quarrelsome--don't sleep of nights--do
not enjoy themselves much, as a general thing;--no, I never go as a
passenger; nor, though I am something of a salt, do I ever go to sea
as a Commodore, or a Captain, or a Cook.  I abandon the glory and
distinction of such offices to those who like them.  For my part, I
abominate all honourable respectable toils, trials, and tribulations
of every kind whatsoever.  It is quite as much as I can do to take
care of myself, without taking care of ships, barques, brigs,
schooners, and what not.  And as for going as cook,--though I confess
there is considerable glory in that, a cook being a sort of officer
on ship-board--yet, somehow, I never fancied broiling fowls;--though
once broiled, judiciously buttered, and judgmatically salted and
peppered, there is no one who will speak more respectfully, not to
say reverentially, of a broiled fowl than I will.  It is out of the
idolatrous dotings of the old Egyptians upon broiled ibis and roasted
river horse, that you see the mummies of those creatures in their
huge bake-houses the pyramids.

No, when I go to sea, I go as a simple sailor, right before the mast,
plumb down into the forecastle, aloft there to the royal mast-head.
True, they rather order me about some, and make me jump from spar to
spar, like a grasshopper in a May meadow.  And at first, this sort of
thing is unpleasant enough.  It touches one's sense of honour,
particularly if you come of an old established family in the land,
the Van Rensselaers, or Randolphs, or Hardicanutes.  And more than
all, if just previous to putting your hand into the tar-pot, you have
been lording it as a country schoolmaster, making the tallest boys
stand in awe of you.  The transition is a keen one, I assure you,
from a schoolmaster to a sailor, and requires a strong decoction of
Seneca and the Stoics to enable you to grin and bear it.  But even
this wears off in time.

What of it, if some old hunks of a sea-captain orders me to get a
broom and sweep down the decks?  What does that indignity amount to,
weighed, I mean, in the scales of the New Testament?  Do you think
the archangel Gabriel thinks anything the less of me, because I
promptly and respectfully obey that old hunks in that particular
instance?  Who ain't a slave?  Tell me that.  Well, then, however the
old sea-captains may order me about--however they may thump and punch
me about, I have the satisfaction of knowing that it is all right;
that everybody else is one way or other served in much the same
way--either in a physical or metaphysical point of view, that is; and
so the universal thump is passed round, and all hands should rub each
other's shoulder-blades, and be content.

Again, I always go to sea as a sailor, because they make a point of
paying me for my trouble, whereas they never pay passengers a single
penny that I ever heard of.  On the contrary, passengers themselves
must pay.  And there is all the difference in the world between
paying and being paid.  The act of paying is perhaps the most
uncomfortable infliction that the two orchard thieves entailed upon
us.  But BEING PAID,--what will compare with it?  The urbane activity
with which a man receives money is really marvellous, considering
that we so earnestly believe money to be the root of all earthly
ills, and that on no account can a monied man enter heaven.  Ah! how
cheerfully we consign ourselves to perdition!

Finally, I always go to sea as a sailor, because of the wholesome
exercise and pure air of the fore-castle deck.  For as in this world,
head winds are far more prevalent than winds from astern (that is, if
you never violate the Pythagorean maxim), so for the most part the
Commodore on the quarter-deck gets his atmosphere at second hand from
the sailors on the forecastle.  He thinks he breathes it first; but
not so.  In much the same way do the commonalty lead their leaders in
many other things, at the same time that the leaders little suspect
it.  But wherefore it was that after having repeatedly smelt the sea
as a merchant sailor, I should now take it into my head to go on a
whaling voyage; this the invisible police officer of the Fates, who
has the constant surveillance of me, and secretly dogs me, and
influences me in some unaccountable way--he can better answer than
any one else.  And, doubtless, my going on this whaling voyage,
formed part of the grand programme of Providence that was drawn up a
long time ago.  It came in as a sort of brief interlude and solo
between more extensive performances.  I take it that this part of the
bill must have run something like this:


"GRAND CONTESTED ELECTION FOR THE PRESIDENCY OF THE UNITED STATES.
"WHALING VOYAGE BY ONE ISHMAEL.
"BLOODY BATTLE IN AFFGHANISTAN."


Though I cannot tell why it was exactly that those stage managers,
the Fates, put me down for this shabby part of a whaling voyage, when
others were set down for magnificent parts in high tragedies, and
short and easy parts in genteel comedies, and jolly parts in
farces--though I cannot tell why this was exactly; yet, now that I
recall all the circumstances, I think I can see a little into the
springs and motives which being cunningly presented to me under
various disguises, induced me to set about performing the part I did,
besides cajoling me into the delusion that it was a choice resulting
from my own unbiased freewill and discriminating judgment.

Chief among these motives was the overwhelming idea of the great
whale himself.  Such a portentous and mysterious monster roused all
my curiosity.  Then the wild and distant seas where he rolled his
island bulk; the undeliverable, nameless perils of the whale; these,
with all the attending marvels of a thousand Patagonian sights and
sounds, helped to sway me to my wish.  With other men, perhaps, such
things would not have been inducements; but as for me, I am tormented
with an everlasting itch for things remote.  I love to sail forbidden
seas, and land on barbarous coasts.  Not ignoring what is good, I am
quick to perceive a horror, and could still be social with it--would
they let me--since it is but well to be on friendly terms with all
the inmates of the place one lodges in.

By reason of these things, then, the whaling voyage was welcome; the
great flood-gates of the wonder-world swung open, and in the wild
conceits that swayed me to my purpose, two and two there floated into
my inmost soul, endless processions of the whale, and, mid most of
them all, one grand hooded phantom, like a snow hill in the air.



CHAPTER 2

The Carpet-Bag.


I stuffed a shirt or two into my old carpet-bag, tucked it under my
arm, and started for Cape Horn and the Pacific.  Quitting the good
city of old Manhatto, I duly arrived in New Bedford.  It was a
Saturday night in December.  Much was I disappointed upon learning
that the little packet for Nantucket had already sailed, and that no
way of reaching that place would offer, till the following Monday.

As most young candidates for the pains and penalties of whaling stop
at this same New Bedford, thence to embark on their voyage, it may as
well be related that I, for one, had no idea of so doing.  For my
mind was made up to sail in no other than a Nantucket craft, because
there was a fine, boisterous something about everything connected
with that famous old island, which amazingly pleased me.  Besides
though New Bedford has of late been gradually monopolising the
business of whaling, and though in this matter poor old Nantucket is
now much behind her, yet Nantucket was her great original--the Tyre
of this Carthage;--the place where the first dead American whale was
stranded.  Where else but from Nantucket did those aboriginal
whalemen, the Red-Men, first sally out in canoes to give chase to the
Leviathan?  And where but from Nantucket, too, did that first
adventurous little sloop put forth, partly laden with imported
cobblestones--so goes the story--to throw at the whales, in order to
discover when they were nigh enough to risk a harpoon from the
bowsprit?

Now having a night, a day, and still another night following before
me in New Bedford, ere I could embark for my destined port, it
became a matter of concernment where I was to eat and sleep
meanwhile.  It was a very dubious-looking, nay, a very dark and
dismal night, bitingly cold and cheerless.  I knew no one in the
place.  With anxious grapnels I had sounded my pocket, and only
brought up a few pieces of silver,--So, wherever you go, Ishmael,
said I to myself, as I stood in the middle of a dreary street
shouldering my bag, and comparing the gloom towards the north with
the darkness towards the south--wherever in your wisdom you may
conclude to lodge for the night, my dear Ishmael, be sure to inquire
the price, and don't be too particular.
*/

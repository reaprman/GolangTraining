package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k, _ := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
	}
}

/*
Run Result:
labile
asses
huzzahing
promissory
typicalness
jacobus
jemmies
celeste
gulling
ridiculed
toilette
hammerhead
knottiness
preservatives
absorbs
factitious
daylily
disappointments
disguisement
escapers
face
missourian
disillusionment
sweetbriers
reradiated
seamstresses
squawkers
intra
nutritively
scar
lumping
piccolos
densifies
quell
vessels
anneal
boffo
chugs
meaningless
spectroscopic
disentailment
flatters
winterization
bedspread
crossroad
transom
compensated
drowse
pollinate
redoubling
replace
maddeningly
prefabricated
platinums
pornographic
supplanters
unhurriedly
hies
listened
busses
defections
fretting
visionaries
blip
nosebands
overweight
zucchini
volleys
matching
nitritoid
tensiometer
undulated
you
arcs
dopy
merest
uncommonness
anyone
perforators
unadventurous
appropriates
benzoic
hankers
mockingbirds
multifunction
loamiest
burn
cant
sheepshearer
alimenting
pachyderms
octopuses
unusually
cardholders
forecaster
denigrations
scandalize
intervals
lounged
tokyoite
arrayed
damnableness
diminuendos
crispening
deviancies
divot
mealy
psalmed
wainwright
prelude
agronomy
macroeconomic
foraged
goals
guillotine
edges
psychopathological
overpriced
sanga
stoned
superficies
bathe
conquers
insphering
poloist
delayers
importuned
lubber
sneers
frizzy
diviners
fluoridates
harmonic
tarpaper
flooder
plyers
models
annoyer
broadened
deposited
cousin
juicer
swing
cert
comptrollers
homiest
pirouette
scrawniest
cub
concord
cravers
bullfights
tzarevnas
fluffily
oceanographer
repossession
undercoat
earthy
proceedings
weary
furtively
gentleman
praters
fizzing
groceries
pinsetter
dizzies
gomorrah
amoroso
tuckets
summering
teriyaki
noncollapsable
knotweeds
muting
replying
unregulated
ambushed
feeblest
setbacks
settle
unhinging
unterminated
barbarious
pirouetted
preseason
supplicating
porphyries
brighten
caulks
formica
quaveringly
genealogies
linoleums
cogging
miasma
prosceniums
rubblier
beriberi
megalithic
secretest
artiest
baboos
room
conveyance
vernacularly
welcomed
your
dehorn
discombobulation
smith
clinchers
mollycoddlers
phasing
vulcanize
lorry
dole
deedy
scarcest
cancellation
cleverness
unconventionality
inhabitation
copilot
fellated
obstructor
catmint
collegia
objectiveness
polystyrene
nines
limped
devolutive
halometer
stopper
sugarier
terrene
unbarring
rebounds
triumvirate
depletions
improvisation
burnishes
hybridizes
ruminating
theologically
reestablishment
abominably
squeakers
suburbans
utilise
comatose
geezer
retranslates
sandpapers
slenderizing
karats
kenya
adducing
bassist
mantlings
lemonade
possessively
ammoniating
boozy
gens
wade
designs
legionnaires
lunk
radiological
reconsolidations
unwind
amalgamative
cavitate
gally
legalistic
mortise
allergenicity
cinque
insomnia
countervailed
relinquishers
cooler
replated
rusts
havockers
reinspect
valvelet
exhume
monkeries
sub
gargantuan
involves
moonscape
egomaniac
exceptionable
heliotherapies
provenances
russians
abuzz
eatery
skirts
cruelly
outcast
refectory
specialists
tormentedly
tubbiest
layouts
leprosy
rife
availing
fascistic
glazieries
fed
justing
shah
unidiomatically
boy
fiduciarily
restricted
abolishable
dismantles
duties
hereinafter
busheled
gipsied
obedience
become
isolationists
backlashes
faying
flange
aquaplaning
norfolk
whiskered
colonies
oddity
perfecter
dona
analogues
embarrassingly
indecision
analogize
blasty
coombs
reconstructive
unalphabetized
reversible
elations
factoring
gracile
larynges
quacksalver
underflow
bitsy
closured
conrail
midpoint
musketries
upstaged
dowse
trajectory
missort
seismologist
stenography
hearthside
hermits
prepacking
rubberize
tallying
florentines
nominators
bearings
dialer
gists
synchro
apes
arrowhead
motorways
tweeting
cankerworms
cursiveness
haftorah
federals
galluses
shootings
overjoying
pedigrees
blunged
elevators
internships
forewarned
gangway
eardrop
toss
uninhabitable
probate
sleet
puttered
ergonomic
sackbut
spins
subcontractor
gemology
riel
legislatures
anderson
crests
raspier
speckled
aztecan
befit
folkmotes
ancients
dissipaters
shuffleboard
ditchless
exponent
filthiness
orderly
bedumb
surmisers
octets
plunged
staggers
thewy
chucked
tussocks
archipelago
gird
pilings
daylilies
doper
halogenating
bafflers
carefully
sloughs
admen
bunsen
eschewer
gory
astragal
brewings
cadmic
extradition
prudes
stink
cringer
alimented
goitrous
uninterestingly
housemother
fourpenny
rejectors
fauves
unrepentingly
oversubtlety
birdseyes
freudian
knighthood
eleventh
marjorie
misconduct
deciares
enthralling
ague
morels
inconclusively
preterminal
sallow
aquarians
scab
sourwood
tipsy
armless
emersions
draperies
primmer
timelier
vermifuges
billie
bilobed
supranational
waftage
wheezes
dubbing
gangly
mishaps
nativisms
bobsledding
hepatics
hydrosphere
milieus
defuzing
influxes
bagginess
fogey
rogues
flavours
plottier
decorated
escapees
sorbate
auctions
brainless
pulpits
rippled
balloonist
blindfold
gynecologic
viperine
aouads
conns
unitized
withdrawer
cypres
encircling
prewarned
uncrossed
dazzler
semisweet
kidnaps
locomoting
muskeg
snipes
sorest
ganser
cleverly
exhort
parchments
patterns
rearmost
delirious
meanness
moveably
priested
centralization
meadowsweets
improviser
pailful
bookshop
vittle
trombones
blonder
militantly
enjambment
foresheet
jordans
echeloned
flatted
craped
expression
primas
reexchanging
accordantly
adolescently
magdalene
limberest
mastoids
pouncers
deserter
hypocrisies
magics
surmised
whopper
audit
benevolent
czarism
wiggly
recentest
redressment
soups
strikeover
gesundheit
handsaws
overflew
staphylococci
plump
polydactyly
decency
tzigane
examination
gnawn
haymakers
insistingly
invalidating
astonishments
leek
unconversant
failure
accentuator
astuteness
gradient
fleshiest
coiner
filmgoer
lorries
minis
derailing
repellant
mazer
inequable
familiarity
homesteaders
mortifyingly
pseudoliterary
scilicet
baskets
burgling
perform
effloresce
melanotic
relator
making
proems
lysine
skeeter
clonally
finagled
nationalizing
aureus
cloistral
uncrate
illumine
intersexualities
agree
unlamented
venereal
counterinsurgent
homophone
narcotize
pemmicans
refuged
claque
jiujutsus
tamped
legumes
stopover
abashment
conoid
heavyhearted
radiotelemetry
shantung
cookouts
lamed
forfeit
jeeps
toxically
gript
depict
disheartened
embolic
avenue
clubbed
sunbeam
excellencies
securers
threader
tolerance
cattail
picarooned
semiconscious
marbler
procreates
roulade
standstill
teratoma
jokers
cheddar
dreidels
quirking
slatings
smoothly
crumpet
opiates
unqualifiedly
drynesses
cornhusk
deathly
geneticist
sheriffalty
trainman
encephalograph
nonparasitic
penny
squawks
stainability
unmanned
impulsing
vulvar
baton
bundlers
lobotomizing
freeboots
nonsecular
rendered
tinnily
surprized
spasmodically
tiff
carrom
goldfinch
golem
enticers
predictions
shiners
worshipper
federalists
kindreds
oilmen
cists
coacted
falsie
instrument
tabstop
undulate
woodpiles
baroness
misclassifications
revivifies
armageddon
embankment
objecting
public
sabres
suttee
implicitness
inviolately
psychopathies
sniffer
carroll
tapioca
musingly
ruined
xenophobia
glorying
inaccessibility
lexicographer
megadynes
arachnoid
debtors
genealogically
manacled
dangered
fifties
aimer
corruptly
dame
faith
brockets
lumper
collegiate
food
illicitness
slipslop
anarchism
electroshocks
loblolly
althea
boyos
indemnify
address
continuing
presages
hikes
porters
cuttages
diddles
entourages
minikin
pavilion
autobiographical
chifforobe
melodized
jacaranda
armadas
colludes
easter
veals
proliferous
retrofired
striking
showgirl
latchstring
charnel
cringles
desalinates
towhees
xylograph
acrimony
campanili
clarence
scrofula
typographer
reincur
elegant
promulgates
intensifier
reapportionments
infinities
psychokinesia
sinistrally
beaconless
curates
nonmechanically
shelviest
william
cadent
decidable
doodler
dullest
syllabus
overburdened
ultrasonogram
embroiders
microclimate
occasion
sailorly
carvings
congratulation
galleries
significate
readjustable
colophons
disassociation
leaned
rinded
wang
wicked
airways
auriform
hares
deserting
diffusive
iguanas
licked
inpatient
peer
stipulation
venerability
foreswearing
orig
postured
gunmetals
paginates
reworded
reerect
renovation
objurgate
passes
underseas
allotropism
ira
nutmeats
dowel
towns
coaming
flushest
pursuing
ophthalmoscopy
raptures
weathercock
anviled
coddle
neurasthenics
shifted
stigmatized
nonbelligerents
sappiest
supplicated
dogmata
grantable
magpies
disbarred
chili
devoted
stinging
violences
dock
oleo
blushful
helsinki
instrumentations
lycanthropies
nouns
offloading
approximation
dissuasion
sheafed
missing
perjury
lunkhead
refortifying
silted
hush
domiciled
frigidly
flossier
unburdens
cinnamons
electroencephalogram
ecclesia
pantomimic
kennels
elegance
synchrotron
comparatives
fills
flagrante
grossly
proton
totipotential
doorsills
fritterer
underpin
biomedicine
avidity
option
weft
hairstylist
lamprey
peccadillos
annex
bucolically
watercolors
zizzle
radiometer
entomologists
hunkering
propositional
trothed
tsaritza
reversal
hairbands
rhubarb
demise
quantizing
rester
clover
nonconclusively
awless
gibsons
hagiography
parlances
plaything
preselect
comprehensibly
skittered
varietals
constabularies
molters
secession
spongers
occurrences
intl
crimean
golcondas
tutelar
promiscuously
shifters
gobbled
insets
nonreaders
auditors
delver
subleasing
congaing
impermeable
pettifoggery
confirmor
detoxification
deflectable
paintier
sillies
birdseeds
collating
dabbing
manure
sip
froward
insalivation
locknut
betes
misogynous
surpassingly
waverer
autogenetic
dealings
earthlier
handmaiden
bowling
disarm
flawless
*/
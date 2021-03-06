package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	// scan the page
	// NewScanner takes a reader and res.Body implements the reader interface (so its a reader)
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// loop over the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

/*
Prins out the word one at a time from the book: this is only a short exert
Run Result:
Some
fell
flat
upon
their
faces.
Like
dislodged
trucks,
the
heads
of
the
harpooneers
aloft
shook
on
their
bull-like
necks.
Through
the
breach,
they
heard
the
waters
pour,
as
mountain
torrents
down
a
flume.
"The
ship!
The
hearse!--the
second
hearse!"
cried
Ahab
from
the
boat;
"its
wood
could
only
be
American!"
Diving
beneath
the
settling
ship,
the
whale
ran
quivering
along
its
keel;
but
turning
under
water,
swiftly
shot
to
the
surface
again,
far
off
the
other
bow,
but
within
a
few
yards
of
Ahab's
boat,
where,
for
a
time,
he
lay
quiescent.
"I
turn
my
body
from
the
sun.
What
ho,
Tashtego!
let
me
hear
thy
hammer.
Oh!
ye
three
unsurrendered
spires
of
mine;
thou
uncracked
keel;
and
only
god-bullied
hull;
thou
firm
deck,
and
haughty
helm,
and
Pole-pointed
prow,--death-glorious
ship!
must
ye
then
perish,
and
without
me?
Am
I
cut
off
from
the
last
fond
pride
of
meanest
shipwrecked
captains?
Oh,
lonely
death
on
lonely
life!
Oh,
now
I
feel
my
topmost
greatness
lies
in
my
topmost
grief.
Ho,
ho!
from
all
your
furthest
bounds,
pour
ye
now
in,
ye
bold
billows
of
my
whole
foregone
life,
and
top
this
one
piled
comber
of
my
death!
Towards
thee
I
roll,
thou
all-destroying
but
*/

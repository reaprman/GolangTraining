package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// an artificial input source
	const input = "Vanity and pride are different things, though the words are often used synonymously. A person may be proud without being vain.Pride relates more to our opinion of ourselves,vanity to what we would have others think of us.Generosity is giving more than you can, andpride is taking less than you need.Anger is a fuel.You need fuel to launch a rocket.But if all you have is fuel without any complex internal mechanism directing it, you don't have a rocket.You have a bomb."

	scanner := bufio.NewScanner(strings.NewReader(input))
	//set the split functon for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input:", err)
	}
}

/*
Run Result:
Vanity
and
pride
are
different
things,
though
the
words
are
often
used
synonymously.
A
person
may
be
proud
without
being
vain.Pride
relates
more
to
our
opinion
of
ourselves,vanity
to
what
we
would
have
others
think
of
us.Generosity
is
giving
more
than
you
can,
andpride
is
taking
less
than
you
need.Anger
is
a
fuel.You
need
fuel
to
launch
a
rocket.But
if
all
you
have
is
fuel
without
any
complex
internal
mechanism
directing
it,
you
don't
have
a
rocket.You
have
a
bomb.
*/

package main

import "fmt"

func main() {

	a := 43
	fmt.Println("a - ", a)
	fmt.Println("a's memory location - ", &a)

	var b *int = &a
	fmt.Println("b - ", b)
	fmt.Println("b's dereferenced - ", *b)

	*b = 42                       // b says, "The value at this address, change to 42"
	fmt.Println("a is now - ", a) // should print 42 for a
}

/*
Run Resuts:
a -  43
a's memory location -  0xc04200c230
b -  0xc04200c230
b's dereferenced -  43
a is now -  42

this is useful
we can pass a memory address instead of a bunch of values
and then we can still change the value of whatever is stored in memory
this makes our program more performant
we don't have to pass around large amounts of data
we only have to pass around addresses

everything is PASS BY VALUE in go
when we pass a memory address, we are passing a value
*/

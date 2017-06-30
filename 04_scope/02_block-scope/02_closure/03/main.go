package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variable used by multiple functions
without closure, for two or more funcs to have access to the same
variable, that variable would need to be package scope

anonymous function
a function without a name: func() int {}

func expresion
assigning a function to a variable

Result
1
2
*/

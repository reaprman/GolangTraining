package main

import "fmt"

func main() {
	age := 44
	changeMe(age)
	fmt.Println(age)
}

func changeMe(z int) {
	fmt.Println(z)
	z = 24
}

/*
When changeMe is called on line 7 the value 44
is being passed as an argument. The variable age's value has
not changed.

Run Results:
44
44
*/

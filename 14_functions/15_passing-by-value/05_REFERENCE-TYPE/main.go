package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) //[Ryan]
}

func changeMe(z []string) {
	z[0] = "Ryan"
	fmt.Println(z) //[Ryan]
}

/*
A slice is a reference type. Its an address pointing to an array.
Run Results:
[]
[Ryan]
[Ryan]
*/

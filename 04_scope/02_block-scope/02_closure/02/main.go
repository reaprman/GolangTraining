package main

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helpsus limit the scope of variables without closure
without closure, for two or more func to have that
variable would need to be package scope

Result:
1
2
*/

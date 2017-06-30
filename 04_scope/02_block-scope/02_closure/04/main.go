package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
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

package main

import "fmt"

func main() {
	func() {
		fmt.Println("I'm driving")
	}() // () makes this a self executing function
}

/*
Run Results:
I'm driving
*/

package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hello World!")
	}
	greeting()
}

/*
Run Results:
Hello World!
*/

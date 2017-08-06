package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hello World!")
	}
	fmt.Printf("%T \n", greeting)
	greeting()
}

/*
The func is anonymous and is being assigned to the variable greeting.
This is called a func expression and is the only way to have a func inside
another function.
Run Results:
func()
Hello World!
*/

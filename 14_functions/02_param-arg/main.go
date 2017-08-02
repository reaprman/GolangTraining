package main

import "fmt"

func main() {
	greet("Jane")
	greet("John")
}

func greet(name string) {
	fmt.Println(name)
}

/*
"greet" is declared with a parameter
When calling greet, pass in an argument

Run Results:
Jane
John
*/

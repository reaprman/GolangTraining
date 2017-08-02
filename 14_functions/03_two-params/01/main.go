package main

import "fmt"

func main() {
	greet("John", "Smith")
}

func greet(fName, lName string) {
	fmt.Println(fName, lName)
}

/*
Run Results:
John Smith
*/

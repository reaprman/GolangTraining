package main

import "fmt"

func main() {

	s := greet("Jane ", "Doe")
	fmt.Println(s)
}

func greet(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

/*
Run Results:

*/

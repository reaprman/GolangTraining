package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 24
	fmt.Printf("%T - %v \n", myAge, myAge)
}

/*
Run Result:
main.foo - 24
*/

package main

import "fmt"

var x int = 42 // scope of x is currently package level

func main() {
	fmt.Println(x)
	foo()
	y := 17 // scope of y is currently block level
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
	// fmt.Println(y) y is not accessible here. block level scope
}

/*
Results

42
42
17
*/

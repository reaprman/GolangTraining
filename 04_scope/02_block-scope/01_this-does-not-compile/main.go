package main

import "fmt"

func main() {
	x := 42
	fmt.Println("x is ", x)
	foo()
}

func foo() {
	fmt.Println("x is ", x) // x defined with block level scope in main.
}

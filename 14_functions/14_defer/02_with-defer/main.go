package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	hello()
}

/*
The defer keyword will defer the running of the function until
the end of the scope of the current function. In this case world()
is defered until before the closing of main().

Run Results:
hello world
*/

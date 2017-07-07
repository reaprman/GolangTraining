package main

import "fmt"

func main() {
	a := 43
	fmt.Println("a - ", a)
	fmt.Println("a's memory location - ", &a)

	var b *int = &a
	fmt.Println("b - ", b)
}

/*
Run Results
a -  43
a's memory location -  0xc0420381c0
b -  0xc0420381c0

the above code makes b a pointer to the memory address where an
int is stored.
b is of type "int pointer"
*int -- the * is part of the type -- b is of type *int
*/

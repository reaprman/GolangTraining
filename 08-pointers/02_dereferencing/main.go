package main

import "fmt"

func main() {
	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory location - ", &a)

	var b *int = &a

	fmt.Println("b - ", b)
	fmt.Println("b's dereference value - ", *b)

}

/*
Run Results:
a -  43
a's memory location -  0xc0420381c0
b -  0xc0420381c0
b's dereference value -  43

b is an int pointer
b points to the memory addrss where an int (a) is stored
to see the value in that memory address, add a * in front of b
the * is an operator in this case
*/

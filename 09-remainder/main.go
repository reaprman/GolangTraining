package main

import "fmt"

func main() {
	x := 13 % 3 // % is used for modulo division and returns the remainder
	fmt.Println(x)
	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	for i := 1; i < 11; i++ {
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}

/*
Run Results:
1
Odd
Odd
Even
Odd
Even
Odd
Even
Odd
Even
Odd
Even

*/

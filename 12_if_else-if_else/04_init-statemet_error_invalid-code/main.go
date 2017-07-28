package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	fmt.Println(food)
}

/*
Code is not executable as error exists on line 13.
the variable food is only accessible within the if statement
it was declared in.
*/

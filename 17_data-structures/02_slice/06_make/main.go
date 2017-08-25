package main

import "fmt"

func main() {
	customerNum := make([]int, 3)
	// 3 is lenght & capacity
	// length - number of elements referred to by the slice
	// capcity - number of elements in the underlying array
	customerNum[0] = 7
	customerNum[1] = 10
	customerNum[2] = 15

	fmt.Println(customerNum[0])
	fmt.Println(customerNum[1])
	fmt.Println(customerNum[2])

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by teh slice
	// 5 is capacity - number of elements in the underlying array
	// you could also do it like this

	greeting[0] = "Good morning"
	greeting[1] = "Bonjour"
	greeting[2] = "dias"

	fmt.Println(greeting[2])

}

/*
Run Result:
7
10
15
dias
*/

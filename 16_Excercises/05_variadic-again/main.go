package main

import "fmt"

func foo(sf ...int) int {
	var total int
	for _, v := range sf {
		total += v
	}
	return total
}

func main() {
	fmt.Println("Total for foo(1,2)", foo(1, 2))
	fmt.Println("Total for foo(1,2,3)", foo(1, 2, 3))
	aSlice := []int{1, 2, 3, 4}
	fmt.Println("Total for foo(aSlice...)", foo(aSlice...))
	fmt.Println("Total for foo()", foo())
}

/*
Run Results
Total for foo(1,2) 3
Total for foo(1,2,3) 6
Total for foo(aSlice...) 10
Total for foo() 0
*/

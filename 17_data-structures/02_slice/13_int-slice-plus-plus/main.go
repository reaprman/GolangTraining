package main

import "fmt"

func main() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++
	// mySlice[0] = mySlice[0] + 1
	// mySlice[0] += 1
	fmt.Println(mySlice[0])
}

/*
Run Result:
0
7
8
*/

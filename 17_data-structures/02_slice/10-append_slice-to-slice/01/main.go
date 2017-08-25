package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	otherSlice := []int{6, 7, 8, 9}

	mySlice = append(mySlice, otherSlice...)

	fmt.Println(mySlice)

}

/*
Run Results:
[1 2 3 4 5 6 7 8 9]
*/

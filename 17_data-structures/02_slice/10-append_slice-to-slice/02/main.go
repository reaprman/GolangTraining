package main

import "fmt"

func main() {

	mySlice := []string{"Monday", "Tuesday"}
	otherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, otherSlice...)
	fmt.Println(mySlice)
}

/*
Run Result:
[Monday Tuesday Wednesday Thursday Friday]
*/

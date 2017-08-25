package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	otherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, otherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)
}

/*
Run Result:
[Monday Tuesday Wednesday Thursday Friday]
[Monday Tuesday Thursday Friday]
*/

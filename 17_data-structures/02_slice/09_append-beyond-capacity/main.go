package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning"
	greeting[1] = "Bonjour"
	greeting[2] = "bueno dias"
	greeting = append(greeting, "Suprabadham")
	greeting = append(greeting, "Zao'an")
	greeting = append(greeting, "Ohayou gazaimasu")
	greeting = append(greeting, "gidday")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}

/*
Capacity of slice has doubled from 5 to 10.
Run Results:
gidday
7
10
*/

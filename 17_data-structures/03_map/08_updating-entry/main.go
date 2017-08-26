package main

import "fmt"

func main() {
	myGreetings := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Bonjour!",
	}

	myGreetings["Harry"] = "Howdy"
	fmt.Println(myGreetings)
	myGreetings["Tim"] = "Hey"
	fmt.Println(myGreetings)

}

/*
Run Result:
map[Harry:Howdy Tim:Good morning Jenny:Bonjour!]
map[Tim:Hey Jenny:Bonjour! Harry:Howdy]
*/

package main

import "fmt"

func main() {
	myGreetings := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Bonjour!",
	}

	myGreetings["Harry"] = "Howdy"
	fmt.Println(len(myGreetings))
	fmt.Println(myGreetings)
}

/*
Run Results:
3
map[Tim:Good morning Jenny:Bonjour! Harry:Howdy]
*/

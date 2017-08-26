package main

import "fmt"

func main() {
	myGreetings := map[string]string{
		"Tim":   "Good Morning",
		"Jenny": "Bonjour!",
	}

	myGreetings["Harry"] = "Howdy"
	fmt.Println(myGreetings)
}

/*
Run Result:
map[Tim:Good Morning Jenny:Bonjour! Harry:Howdy]
*/

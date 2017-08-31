package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		1: "Good morning.",
		2: "Bonjour",
		3: "Hola!",
		4: "Bongiorno",
	}
	myGreeting[5] = "Ohayo!"

	fmt.Println(myGreeting)
	delete(myGreeting, 7)
	fmt.Println(myGreeting)
}

/*
There is no element with key 7 in the map. The delete,
however allows the code to run with no errors.

Run Result:
map[5:Ohayo! 1:Good morning. 2:Bonjour 3:Hola! 4:Bongiorno]
map[1:Good morning. 2:Bonjour 3:Hola! 4:Bongiorno 5:Ohayo!]
*/

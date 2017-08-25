package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)

	myGreeting["Todd"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting == nil)
	fmt.Println(myGreeting)
}

/*
Run Result:
false
map[Todd:Good morning. Jenny:Bonjour.]
*/

package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Have a great day!",
	}

	fmt.Println(myGreeting == nil)
	fmt.Println(myGreeting)
}

/*
Run Result:
false
map[Tim:Good morning! Jenny:Have a great day!]
*/

package main

import "fmt"

func main() {
	myGreetings := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "Buenos dias",
		3: "Bongiorno!",
	}
	fmt.Println(myGreetings)
	delete(myGreetings, 2)
	fmt.Println(myGreetings)

}

/*
Run Result:
map[3:Bongiorno! 0:Good morning 1:Bonjour 2:Buenos dias]
map[0:Good morning 1:Bonjour 3:Bongiorno!]
*/

package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "Buenos dias",
		3: "Bongiorno!",
	}
	fmt.Println(myGreeting)

	if val, exists := myGreeting[2]; exists {
		delete(myGreeting, 2)
		fmt.Println("Val: ", val)
		fmt.Println("exists?: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists?: ", exists)
	}
	fmt.Println(myGreeting)
}

/*
Run Results:
map[2:Buenos dias 3:Bongiorno! 0:Good morning 1:Bonjour]
Val:  Buenos dias
exists?:  true
map[3:Bongiorno! 0:Good morning 1:Bonjour]
*/

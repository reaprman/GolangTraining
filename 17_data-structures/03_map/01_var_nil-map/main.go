package main

import "fmt"

func main() {
	var myGreeting map[string]string
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}

/*
The code above creates a nil reference NOT THE CORRECT WAY TO USE MAP

Run Result:
panic: assignment to entry in nil map

goroutine 1 [running]:
main.main()
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTra
ining/17_data-structures/03_map/01_var_nil-map/main.go:7 +0x60
exit status 2
*/

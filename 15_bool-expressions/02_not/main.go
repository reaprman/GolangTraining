package main

import "fmt"

func main() {
	if !true {
		fmt.Println("This did not run")
	}
	if !false {
		fmt.Println("This did run")
	}
}

/*
not operator - !

Run Result:
This did run
*/

package main

import "fmt"

func main() {
	if !true { // !true - NOT true
		fmt.Println("This did not run")
	}
	if !false { // !false - NOT false
		fmt.Println("This did run")
	}
}

/*

! : means NOT in Golang. So !true is the same as saying false

Run Result:
This did run
*/

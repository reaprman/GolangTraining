package main

import "fmt"

func main() {
	name := "Ryan"
	fmt.Println(name) // Ryan

	changeMe(name)
	fmt.Println(name) //Ryan

}

func changeMe(z string) {
	fmt.Println(z) // Ryan
	z = "John-117"
	fmt.Println(z) // John-117
}

/*
Run Result:
Ryan
Ryan
John-117
Ryan
*/

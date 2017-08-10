package main

import "fmt"

func main() {
	name := "Ryan"
	fmt.Println(name) //Ryan

	changeMe(&name)
	fmt.Println(&name) //0xc04203a1c0
	fmt.Println(name)  //John-117
}

func changeMe(z *string) {
	fmt.Println(z)  //0xc04203a1c0
	fmt.Println(*z) //Ryan
	*z = "John-117"
	fmt.Println(z)  //0xc04203a1c0
	fmt.Println(*z) //John-117
}

/*
Run Results:
Ryan
0xc04203a1c0
Ryan
0xc04203a1c0
John-117
0xc04203a1c0
John-117
*/

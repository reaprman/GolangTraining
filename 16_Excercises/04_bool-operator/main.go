package main

import "fmt"

func main() {
	fmt.Println((true && false) || (false && true) || !(false && false))
}

/*
Whats the value of the following expressions:
(true && false) || (false && true) || !(false && false)

TRUE

*/

package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T - %v \n", myAge, myAge)

	var yourAge int
	yourAge = 27
	fmt.Printf("%T - %v \n", yourAge, yourAge)

	//this doesn't work
	//fmt.Println(myAge + yourAge)

	//this conversion works
	fmt.Println(int(myAge) + yourAge)
}

/*
1st Run Result:
invalid operation: myAge + yourAge (mismatched types foo and i
nt)

2nd Run Result:
main.foo - 44
int - 27
71
*/

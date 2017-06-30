package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The credit belongs with the one who is in the ring."
		fmt.Println(y)
	}
	//fmt.Println(y) // outside scope of y
}

/*
First run error:
# command-line-arguments
.\main.go:13: undefined: y

Second run:
42
42
The credit belongs with the one who is in the ring.
*/

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
	//fmt.Println(y)
}

/*
Error occurs on second run due to y being defined in the inner scope of
func main(). So it is not accessible in the next Println

1st Run Results:
42
42
The credit belongs with the one who is in the ring.

2nd Run Results:
.\main.go:13: undefined: y
*/

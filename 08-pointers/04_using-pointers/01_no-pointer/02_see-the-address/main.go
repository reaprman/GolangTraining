package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", z) // address in func zero
	fmt.Println(&z)       // address in func zero
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", x) // address in func zero
	fmt.Println(&x)       // address in func zero
	zero(x)
	fmt.Println(x) // x is still 5
}

/*
Run Results:
%!p(int=5)
0xc04200c230
%!p(int=5)
0xc04200c248
5

the value is passed with the statment: zero(x)

*/

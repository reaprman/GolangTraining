package main

import "fmt"

func main() {

	age := 44
	fmt.Println(&age) //0xc04200c230

	changeMe(&age)

	fmt.Println(&age) //0xc04200c230
	fmt.Println(age)  //24
}

func changeMe(z *int) {
	fmt.Println(z)  //0xc04200c230
	fmt.Println(*z) //4
	*z = 24
	fmt.Println(z)  //0xc04200c230
	fmt.Println(*z) //24
}

/*
When changeMe is called on line 10
the value 44 is being passed as an argument

Run Results:
0xc04200c230
0xc04200c230
44
0xc04200c230
24
0xc04200c230
24
*/

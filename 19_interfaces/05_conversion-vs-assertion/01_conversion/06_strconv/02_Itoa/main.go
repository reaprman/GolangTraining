package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 12
	y := "I have this many: " + strconv.Itoa(x)
	fmt.Println(y)
//	fmt.Println("I have this many: ", strconv.Itoa(x), x)
}

/*
1st Run Result:
I have this many: 12

2nd:
I have this many:  12 12
*/
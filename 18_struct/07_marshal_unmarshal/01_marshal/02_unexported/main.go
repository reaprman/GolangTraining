package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}

/*

There are no values after the marshal because the variables for
the Person struct are not exported as they start with lower cases
so no values are but into the []byte(bs).

Run Result:
{James Bond 20}
[]uint8
{}
*/

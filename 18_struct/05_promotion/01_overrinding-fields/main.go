package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   25,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}
	p2 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}
	// fields and methods of an inner type are promoted to the outer type
	fmt.Println(p1.First, p1.Person.First)
	fmt.Println(p2.First, p2.Person.First)
}

/*
Run Result:
Double Zero Seven James
If looks could kill James
*/

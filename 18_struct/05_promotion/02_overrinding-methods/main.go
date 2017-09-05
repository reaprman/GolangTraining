package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

// Greeting - struct examples

func (p Person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz DoubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see again")
}

func main() {
	p1 := Person{
		Name: "Tony Stark",
		Age:  44,
	}
	p2 := DoubleZero{
		Person: Person{
			Name: "James Bond",
			Age:  25,
		},
		LicenseToKill: true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
}

/*
Run Result:
I'm just a regular person.
Miss Moneypenny, so good to see again
I'm just a regular person.
*/

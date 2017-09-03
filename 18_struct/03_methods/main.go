package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + p.last
}

func main() {
	p1 := person{"James", "Bond", 25}
	p2 := person{"Miss", "Moneypenny", 21}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}

/*
Run Result:
amesBond
MissMoneypenny
*/

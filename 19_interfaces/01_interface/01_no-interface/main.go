package main

import "fmt"

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func main() {
	s := square{10}
	fmt.Println("Area: ", s.area())
}

/*
Run Result:
Area:  100
*/

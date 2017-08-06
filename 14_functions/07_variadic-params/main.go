package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64 // total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

/*
you can youse range to iterate a list of objects or collection of items
it returns the iteration number and the current value being used.

Run Results:
[43 56 87 12 45 57]
[]float64
50
*/

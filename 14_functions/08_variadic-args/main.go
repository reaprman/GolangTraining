package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57} // slice of float64's
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

/*
the ... used with the variable data allows us to take this slice
of float64s and pass them one at a time to match the average func
declaration. Alternatively we could just use []float64 as the type
average takes as a parameter.

Run Results:
[43 56 87 12 45 57]
[]float64
50
*/

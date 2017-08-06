package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data)
	fmt.Println(n)
}

func average(sf []float64) float64 {
	total := 0.0
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

/*
data is declared as an array of float64 and passed to a function
to obtain the average of the numbers in the array
Run Result
[43 56 87 12 45 57]
[]float64
50
*/
package main

import "fmt"

func main() {

	var nSum int //  nSum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			nSum += i
		} else if i%5 == 0 {
			nSum += i
		}
	}
	fmt.Println("Sum: ", nSum)
}

/*
If we list all the natural numbers below 10 that are multiples
of 3 or 5, we get 3,5,6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5, below 1000

Run Results:
Sum:  233168
*/

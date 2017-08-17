package main

import "fmt"

func main() {
	evaluateMe := func(x int) (float64, bool) {
		if x%2 == 0 {
			return float64(x / 2), true
		}
		return float64(x / 2), false
	}

	fmt.Println(evaluateMe(5))
}

/*
Run Result:
2 false
*/

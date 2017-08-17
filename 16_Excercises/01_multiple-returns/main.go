package main

import "fmt"

func evaluateMe(x int) (float64, bool) {
	if x%2 == 0 {
		return (float64(x) / 2), true
	}

	return (float64(x) / 2), false
}

func main() {
	fmt.Println(evaluateMe(9))
}

/*
1st Run Results: passing 10 by value
5 True

2nd Run Results: passing 9 by value
4.5 false
*/

package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	ar := []int{}
	for _, n := range numbers {
		if callback(n) {
			ar = append(ar, n)
		}
	}
	return ar
}

func main() {
	ar := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(ar) //[2 3 4]
}

/*
Run Result
[2 3 4]
*/

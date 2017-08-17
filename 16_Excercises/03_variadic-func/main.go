package main

import "fmt"

func main() {
	largest := calculateMe(3, 2, 5, 6, 4400, 6, 76, 586, 958, 984)
	fmt.Println(largest)
}

func calculateMe(nlist ...int) int {
	var great int
	for _, v := range nlist {
		if v > great {
			great = v
		}
	}
	return great
}

/*
Run Result
4400
*/

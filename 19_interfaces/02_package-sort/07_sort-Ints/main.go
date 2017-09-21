package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)
}
/*
Run Result:
[7 4 8 2 9 19 12 32 3]
[2 3 4 7 8 9 12 19 32]
*/
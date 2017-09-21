package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}

/*
Run Result:
[7 4 8 2 9 19 12 32 3]
[32 19 12 9 8 7 4 3 2]
*/
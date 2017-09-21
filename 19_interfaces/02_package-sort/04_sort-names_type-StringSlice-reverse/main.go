package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}

/*
Run Result:
[Zeno John Al Jenny]
[Zeno John Jenny Al]
*/


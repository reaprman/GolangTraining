package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.StringSlice(s).Sort()
	//sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}

/*

1st Run Result:
[Zeno John Al Jenny]
[Al Jenny John Zeno]

2nd Run Result:
[Zeno John Al Jenny]
[Al Jenny John Zeno]

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
}

/*
Run Result:
[Zeno John Al Jenny]
[Al Jenny John Zeno]
12640
*/
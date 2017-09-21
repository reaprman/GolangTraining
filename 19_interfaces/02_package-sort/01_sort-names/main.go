package main

import (
	"fmt"
	"sort"
)

type People []string

func (p People) Len() int { return len(p) }
func (p People) Less(i, j int) bool { return p[i] < p[j] }
func (p People) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main(){

	studyGroup := People{"Zeno", "John", "Al", "Jenny"}
	fmt.Printf("%T\n", studyGroup)
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}

/*
Run Result:
main.People
[Zeno John Al Jenny]
[Al Jenny John Zeno]
*/
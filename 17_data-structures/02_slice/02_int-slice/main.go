package main

import "fmt"

func main() {

	myslice := []int{1, 3, 5, 7, 9, 11}

	for i, value := range myslice {
		fmt.Println(i, " - ", value)
	}
}

/*
Run Results:
0  -  1
1  -  3
2  -  5
3  -  7
4  -  9
5  -  11
*/

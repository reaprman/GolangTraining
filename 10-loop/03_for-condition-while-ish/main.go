package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println("i - ", i)
		i++
	}
}

/*
Run Results:

i -  0
i -  1
i -  2
i -  3
i -  4
i -  5
i -  6
i -  7
i -  8
i -  9
*/

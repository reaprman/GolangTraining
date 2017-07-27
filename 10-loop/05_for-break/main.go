package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println("i -", i)
		if i >= 10 {
			break
		}
		i++
	}
}

/*
Run Results:
i - 0
i - 1
i - 2
i - 3
i - 4
i - 5
i - 6
i - 7
i - 8
i - 9
i - 10
*/

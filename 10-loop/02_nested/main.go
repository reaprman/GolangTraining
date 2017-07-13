package main

import "fmt"

func main() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Println("i - ", i, " : j - ", j)
		} // end of inner loop
	} // end of outer loop
}

/*
Run Results:
i -  0  : j -  0
i -  0  : j -  1
i -  0  : j -  2
i -  0  : j -  3
i -  0  : j -  4
i -  0  : j -  5
i -  1  : j -  0
i -  1  : j -  1
i -  1  : j -  2
i -  1  : j -  3
i -  1  : j -  4
i -  1  : j -  5
i -  2  : j -  0
i -  2  : j -  1
i -  2  : j -  2
i -  2  : j -  3
i -  2  : j -  4
i -  2  : j -  5
i -  3  : j -  0
i -  3  : j -  1
i -  3  : j -  2
i -  3  : j -  3
i -  3  : j -  4
i -  3  : j -  5
i -  4  : j -  0
i -  4  : j -  1
i -  4  : j -  2
i -  4  : j -  3
i -  4  : j -  4
i -  4  : j -  5
i -  5  : j -  0
i -  5  : j -  1
i -  5  : j -  2
i -  5  : j -  3
i -  5  : j -  4
i -  5  : j -  5

.....
*/

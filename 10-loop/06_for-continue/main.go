package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println("i - ", i)
		if i >= 50 {
			break
		}
	}
}

/*
Program prints all odd numbers from 1-51

Run Results:
i -  1
i -  3
i -  5
i -  7
i -  9
i -  11
i -  13
i -  15
i -  17
i -  19
i -  21
i -  23
i -  25
i -  27
i -  29
i -  31
i -  33
i -  35
i -  37
i -  39
i -  41
i -  43
i -  45
i -  47
i -  49
i -  51
*/

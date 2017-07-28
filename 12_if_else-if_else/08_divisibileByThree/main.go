package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("i - ", i)
		}
	}
}

/*
Returns the numbers that are
divisible by 3 as they have no remainder

Run Result:

i -  0
i -  3
i -  6
i -  9
i -  12
i -  15
i -  18
i -  21
i -  24
i -  27
i -  30
i -  33
i -  36
i -  39
i -  42
i -  45
i -  48
i -  51
i -  54
i -  57
i -  60
i -  63
i -  66
i -  69
i -  72
i -  75
i -  78
i -  81
i -  84
i -  87
i -  90
i -  93
i -  96
i -  99
*/

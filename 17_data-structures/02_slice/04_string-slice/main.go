package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning",
		"Bonjour",
		"Dias",
		"Bongiorno",
		"Ohayo",
		"Selamet pagi",
		"Gutten morgen",
	}

	for i, currEntry := range greeting {
		fmt.Println(i, " - ", currEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}
}

/*
Run Result:
0  -  Good morning
1  -  Bonjour
2  -  Dias
3  -  Bongiorno
4  -  Ohayo
5  -  Selamet pagi
6  -  Gutten morgen
Good morning
Bonjour
Dias
Bongiorno
Ohayo
Selamet pagi
Gutten morgen
*/

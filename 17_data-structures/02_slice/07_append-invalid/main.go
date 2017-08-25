package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in thte underlying array

	greeting[0] = "Good morning"
	greeting[1] = "Bonjour"
	greeting[2] = "bueno dias"
	greeting[3] = "suprabadham"

	fmt.Println(greeting[3])

}

/*
Code causes error on line 13. This is the incorrect to append an additional item
to a slice after ti has reached its max length.

Run Results:
panic: runtime error: index out of range

goroutine 1 [running]:
main.main()
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTra
ining/17_data-structures/02_slice/07_append-invalid/main.go:13 +0x6f
exit status 2
*/

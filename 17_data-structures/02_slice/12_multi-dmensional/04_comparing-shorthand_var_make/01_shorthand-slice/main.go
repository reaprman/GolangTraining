package main

import "fmt"

func main() {
	student := []string{}
	students := [][]string{}
	student[0] = "Todd"
	//student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}

/*
Underlying array has not been specified

1st Run Result:
panic: runtime error: index out of range

goroutine 1 [running]:
main.main()
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTra
ining/17_data-structures/02_slice/12_multi-dmensional/04_comparing-shorthand_var_make/
main.go:8 +0x49
exit status 2

2nd Run Result:
[Todd]
[]
*/

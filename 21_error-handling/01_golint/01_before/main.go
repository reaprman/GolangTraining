package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(n int) string {
	if n > 10 {
		return fmt.Sprintf("x is greater than 10")
	} else {
		return fmt.Sprint("x is less than 10")
	}
}

/*
golint result:
main.go:14:9: if block ends with a return statement, so drop this else and outdent its
block

*/
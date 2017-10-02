package main

import "fmt"

func main() {
	var val interface{} = 7
	fmt.Println(val + 6)
}

/*
Run Result:
.\main.go:7: invalid operation: val + 6 (mismatched types interface {} and in
t)
*/

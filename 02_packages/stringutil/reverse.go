// Package stringutil contains utility function for working with
package stringutil

// Reverse returns its argument sring reversed rune-wise left to right
func Reverse(s string) string {
	return reverseTwo(s)
}

/*
go build
	go build reverse.go reverseTwo.go
	won't produce an output file.

go install
	will place package
*/

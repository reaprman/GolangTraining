package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

// Sqrt - function to perform square root
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return 42, nil
}

/*
see use of errors.New in standard library:
http://golang.org/src/pkg/bufio/bufio.go
http://golang.org/src/pkg/io/io.go

Run Result
$ go run main.go
*errors.errorString
2017/11/03 01:07:30 norgate math: square root of negative number
exit status 1
*/

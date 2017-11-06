package main

import (
	"errors"
	"log"
)

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

// Sqrt - function to perform square root
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	//implmentation
	return 42, nil
}

/*
Run Result:
2017/11/02 23:43:40 norgate math: square root of negative number
exit status 1
*/

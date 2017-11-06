package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := Sqrt(-15.7)
	if err != nil {
		log.Fatalln(err)
	}
}

// Sqrt - function to perform square root
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		ErrNorgateMath := fmt.Errorf("norgate math again: square root of negative number: %v", f)
		return 0, ErrNorgateMath
	}
	//implementation
	return 42, nil
}

/*
Run Result:
2017/11/03 01:44:40 norgate math again: square root of negative number: -15.7
exit status 1
*/

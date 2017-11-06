package main

import (
	"fmt"
	"log"
)

type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occurred: %v %v %v",
		n.lat, n.long, n.err)
}

func main() {
	_, err := Sqrt(-15.7)
	if err != nil {
		log.Fatalln(err)
	}
}

// Sqrt - function to perform square root
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, &NorgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}

/*
Run Result:
2017/11/03 01:58:50 a norgate math error occurred: 50.2289 N 99.4656 W
norgate math redux: square root of negative number: -15.7
exit status 1
*/

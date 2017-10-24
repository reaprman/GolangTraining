package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}

/*
Run Result:
1
*/

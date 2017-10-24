package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println("Channel: ", <-c)
		}
	}()
	time.Sleep(time.Second)
}

/*
Run Result:
Channel:  0
Channel:  1
Channel:  2
Channel:  3
Channel:  4
Channel:  5
Channel:  6
Channel:  7
Channel:  8
Channel:  9
*/

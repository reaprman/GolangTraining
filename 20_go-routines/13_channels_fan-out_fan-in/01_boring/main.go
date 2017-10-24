package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring: I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// Fan IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1  // take value off of input1 and put it on chan c
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

/*
Run Result:
Joe 0
Ann 0
Ann 1
Joe 1
Ann 2
Joe 2
Ann 3
Joe 3
Ann 4
Joe 4
You're both boring: I'm leaving.
*/

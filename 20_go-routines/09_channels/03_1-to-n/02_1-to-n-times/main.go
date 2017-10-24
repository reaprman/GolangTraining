package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)
	n := 5

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()
	}

	// this holds main until the above processes run/finish
	for i := 0; i < n; i++ {
		<-done
	}
}

/*
Run Result:
0
1
4
5
7
9
3
2
8
6
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions () that all read from in.
	xc := fanOut(in, 10)

	//TROUBLESHOOTING
	fmt.Printf("%T \n", xc)
	fmt.Println("****************************", len(xc))
	for _, v := range xc {
		fmt.Println("*********", <-v)
	}

	// FAN IN
	// Multiplex multiple channels onto a single channel
	// merge the channels from   onto a single channel
	for n := range merge(xc...) {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int, n int) []<-chan int {
	//xc := make([]<-chan int, n)
	var xc []<-chan int // this needed to be zero?
	for i := 0; i < n; i++ {
		xc = append(xc, factorial(in))
	}
	return xc
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done. This must start after the wg.Add call.

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
Run Result:
[]<-chan int
**************************** 10
********* 720
********* 120
********* 24
********* 6
********* 40320
********* 362880
********* 5040
********* 3628800
********* 39916800
********* 479001600
6
24
5040
720
3628800
40320
120
362880
479001600
39916800
6
24
120
720
5040
40320
479001600
362880
3628800
39916800
6
24
120
720
5040
40320
362880
3628800
39916800
479001600
6
24
120
720
5040
40320
362880
3628800
39916800
479001600
6
24
120
720
5040
40320
362880
3628800
39916800
479001600
6
24
120
720
5040
40320
362880
3628800
39916800
479001600
6
24
120
720
5040
40320
362880
3628800
39916800
479001600
6
24
120
720
5040
40320
362880
3628800
39916800
479001600
6
24
120
720
5040
40320
362880
3628800
39916800
479001600
*/
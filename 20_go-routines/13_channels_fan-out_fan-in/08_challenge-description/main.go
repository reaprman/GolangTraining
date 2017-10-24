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

func fanOut(in <- chan int, n int) []<-chan int {
	xc := make([]<-chan int, n)
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
6
120
720
24
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
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:19
+0x8e

goroutine 29 [chan receive (nil chan)]:
main.merge.func1(0x0)
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:69
+0x5e
created by main.merge
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:76
+0x10d

goroutine 30 [chan receive (nil chan)]:
main.merge.func1(0x0)
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:69
+0x5e
created by main.merge
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:76
+0x10d

goroutine 31 [chan receive (nil chan)]:
main.merge.func1(0x0)
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:69
+0x5e
created by main.merge
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:76
+0x10d

goroutine 32 [chan receive (nil chan)]:
main.merge.func1(0x0)
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:69
+0x5e
created by main.merge
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:76
+0x10d

goroutine 33 [chan receive (nil chan)]:
main.merge.func1(0x0)
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:69
+0x5e
created by main.merge
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:76
+0x10d

goroutine 34 [chan receive (nil chan)]:
main.merge.func1(0x0)
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:69
+0x5e
created by main.merge
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:76
+0x10d

goroutine 35 [chan receive (nil chan)]:
main.merge.func1(0x0)
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:69
+0x5e
created by main.merge
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:76
+0x10d

goroutine 36 [chan receive (nil chan)]:
main.merge.func1(0x0)
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:69
+0x5e
created by main.merge
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:76
+0x10d

goroutine 37 [chan receive (nil chan)]:
main.merge.func1(0x0)
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:69
+0x5e
created by main.merge
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:76
+0x10d

goroutine 38 [chan receive (nil chan)]:
main.merge.func1(0x0)
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:69
+0x5e
created by main.merge
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:76
+0x10d

goroutine 49 [semacquire]:
sync.runtime_Semacquire(0xc0420381cc)
        C:/Go/src/runtime/sema.go:47 +0x3b
sync.(*WaitGroup).Wait(0xc0420381c0)
        C:/Go/src/sync/waitgroup.go:131 +0x81
main.merge.func2(0xc0420381c0, 0xc0420364e0)
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:83
+0x32
created by main.merge
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTr
aining/20_go-routines/13_channels_fan-out_fan-in/08_challenge-description/main.go:85
+0x14c
exit status 2
*/
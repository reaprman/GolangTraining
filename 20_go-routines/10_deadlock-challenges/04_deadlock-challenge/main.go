package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for {
		fmt.Println(<-c)
	}
}

/*
This prints the number, but we have received this error:
fatal error: all goroutines are asleep - deadlock!
Where is the deadlock?
Why are all goroutines asleep
How can we fix this

Run Result:
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangT
raining/20_go-routines/10_deadlock-challenges/04_deadlock-challenge/main.go:15 +0x95

exit status 2
*/

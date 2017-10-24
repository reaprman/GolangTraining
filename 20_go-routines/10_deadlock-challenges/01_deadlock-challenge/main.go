package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}
/*
This results in a deadlock.
Can you determine why?
And what would you do to fix it?

Run Result:
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangT
raining/20_go-routines/10_deadlock-challenges/01_deadlock-challenge/main.go:7 +0x73
exit status 2
*/
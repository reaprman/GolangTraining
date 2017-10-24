package main

import "fmt"

func main() {
	c := factorial(4)
	//fmt.Println("Total: ",<-c)
	for n := range c {
		fmt.Println("Total: ", n)
	}

}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

/*
Run Result:
Total:  24

If it was required to run many factorial calculations at once then
this could be a reason to use go routines to perform the calculation
or if the program ask the user for what factorials to make and the number was large
it could do the processing and then still start processing the next number by asking
the user for the next factorial while the current one is running in another go routine.
*/
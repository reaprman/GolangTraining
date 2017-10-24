package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

 	for n := range c {
 		fmt.Println(n)
 	}
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-c)
	// }
}

/*
If you range you must close the channel otherwise you can loop
the same number of values on the channel

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
*/

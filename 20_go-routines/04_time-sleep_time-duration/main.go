package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

/*
Run Result:
Bar:  0
Foo:  0
Foo:  1
Bar:  1
Bar:  2
Foo:  2
Foo:  3
Bar:  3
Bar:  4
Foo:  4
Foo:  5
Bar:  5
Bar:  6
Foo:  6
Foo:  7
Bar:  7
Bar:  8
Foo:  8
Foo:  9
Bar:  9
*/

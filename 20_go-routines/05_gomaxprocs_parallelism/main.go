package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 15; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 15; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}

/*
Run Result:
Bar:  0
Foo:  0
Foo:  1
Foo:  2
Bar:  1
Foo:  3
Bar:  2
Foo:  4
Foo:  5
Bar:  3
Foo:  6
Foo:  7
Bar:  4
Foo:  8
Foo:  9
Bar:  5
Foo:  10
Bar:  6
Foo:  11
Foo:  12
Bar:  7
Foo:  13
Bar:  8
Foo:  14
Bar:  9
Bar:  10
Bar:  11
Bar:  12
Bar:  13
Bar:  14
*/

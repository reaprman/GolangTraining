package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incrementor("Foo: ")
	go incrementor("Bar: ")
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		//race:
		//counter++
		//no race:
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter: ", counter)
	}
	wg.Done()
}

/*
Run Result:
Bar:  3 Counter:  4
Foo:  0 Counter:  5
Bar:  4 Counter:  6
Foo:  1 Counter:  7
Foo:  2 Counter:  8
Foo:  3 Counter:  10
Bar:  5 Counter:  9
Bar:  6 Counter:  11
Bar:  7 Counter:  12
Foo:  4 Counter:  13
Bar:  8 Counter:  14
Foo:  5 Counter:  15
Foo:  6 Counter:  16
Foo:  7 Counter:  17
Bar:  9 Counter:  18
Bar:  10 Counter:  19
Foo:  8 Counter:  20
Foo:  9 Counter:  21
Bar:  11 Counter:  22
Bar:  12 Counter:  24
Bar:  13 Counter:  25
Foo:  10 Counter:  24
Foo:  11 Counter:  26
Bar:  14 Counter:  27
Bar:  15 Counter:  28
Bar:  16 Counter:  29
Foo:  12 Counter:  30
Foo:  13 Counter:  31
Bar:  17 Counter:  32
Foo:  14 Counter:  33
Foo:  15 Counter:  34
Bar:  18 Counter:  35
Bar:  19 Counter:  36
Foo:  16 Counter:  37
Foo:  17 Counter:  38
Foo:  18 Counter:  39
Foo:  19 Counter:  40
Final Counter:  40
*/

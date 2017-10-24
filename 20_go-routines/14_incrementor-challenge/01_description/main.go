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
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity

Run Result:
Foo:  0 Counter:  1
Foo:  1 Counter:  2
Bar:  0 Counter:  3
Bar:  1 Counter:  4
Bar:  2 Counter:  5
Bar:  3 Counter:  6
Foo:  2 Counter:  7
Foo:  3 Counter:  8
Foo:  4 Counter:  9
Bar:  4 Counter:  10
Bar:  5 Counter:  11
Bar:  6 Counter:  13
Foo:  5 Counter:  12
Foo:  6 Counter:  14
Bar:  7 Counter:  15
Bar:  8 Counter:  16
Foo:  7 Counter:  17
Bar:  9 Counter:  18
Foo:  8 Counter:  19
Bar:  10 Counter:  20
Bar:  11 Counter:  21
Foo:  9 Counter:  22
Foo:  10 Counter:  23
Foo:  11 Counter:  25
Bar:  12 Counter:  24
Bar:  13 Counter:  26
Foo:  12 Counter:  27
Foo:  13 Counter:  28
Foo:  14 Counter:  30
Bar:  14 Counter:  29
Bar:  15 Counter:  31
Foo:  15 Counter:  32
Bar:  16 Counter:  33
Bar:  17 Counter:  34
Foo:  16 Counter:  35
Foo:  17 Counter:  36
Foo:  18 Counter:  37
Foo:  19 Counter:  39
Bar:  18 Counter:  38
Bar:  19 Counter:  40
Final Counter:  40
*/
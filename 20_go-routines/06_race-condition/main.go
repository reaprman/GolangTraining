package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementor("Foo: ")
	go incrementor("Bar: ")
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter: ", counter)
	}
	wg.Done()
}

/*
go run main.go
go run -race main.go
Run Result:
Bar:  0 Counter:  1
Bar:  1 Counter:  2
==================
WARNING: DATA RACE
Write at 0x000000f5a2c0 by goroutine 6:
  main.incrementor()
      C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/Go
langTraining/20_go-routines/06_race-condition/main.go:26 +0xab

Previous write at 0x000000f5a2c0 by goroutine 7:
  main.incrementor()
      C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/Go
langTraining/20_go-routines/06_race-condition/main.go:26 +0xab

Goroutine 6 (running) created at:
  main.main()
      C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/Go
langTraining/20_go-routines/06_race-condition/main.go:15 +0x7b

Goroutine 7 (running) created at:
  main.main()
      C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/Go
langTraining/20_go-routines/06_race-condition/main.go:16 +0xa8
==================
Foo:  0 Counter:  1
Bar:  2 Counter:  3
Foo:  1 Counter:  2
Bar:  3 Counter:  3
Foo:  2 Counter:  4
Bar:  4 Counter:  4
Bar:  5 Counter:  5
Foo:  3 Counter:  5
Foo:  4 Counter:  6
Bar:  6 Counter:  6
Foo:  5 Counter:  7
Foo:  6 Counter:  8
Bar:  7 Counter:  8
Bar:  8 Counter:  9
Foo:  7 Counter:  9
Bar:  9 Counter:  10
Foo:  8 Counter:  10
Bar:  10 Counter:  11
Bar:  11 Counter:  12
Foo:  9 Counter:  11
Foo:  10 Counter:  13
Foo:  11 Counter:  14
Foo:  12 Counter:  15
Bar:  12 Counter:  13
Foo:  13 Counter:  14
Foo:  14 Counter:  15
Foo:  15 Counter:  16
Bar:  13 Counter:  14
Bar:  14 Counter:  15
Foo:  16 Counter:  17
Bar:  15 Counter:  16
Bar:  16 Counter:  17
Foo:  17 Counter:  18
Foo:  18 Counter:  19
Foo:  19 Counter:  20
Bar:  17 Counter:  18
Bar:  18 Counter:  19
Bar:  19 Counter:  20
Final Counter:  20
Found 1 data race(s)
exit status 66
*/
package main

import "fmt"

func main() {
	foo()
	bar()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo: ", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar: ", i)
	}
}

/*
No concurrency happens

Run Result:
Foo:  0
Foo:  1
Foo:  2
Foo:  3
Foo:  4
Foo:  5
Foo:  6
Foo:  7
Foo:  8
Foo:  9
Bar:  0
Bar:  1
Bar:  2
Bar:  3
Bar:  4
Bar:  5
Bar:  6
Bar:  7
Bar:  8
Bar:  9
*/
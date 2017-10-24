package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// 	fmt.Println("err happened:", err)
		//	log.Println("err happened:", err)
		//	log.Fatalln(err)
		panic(err)
	}
}

/*
Package log implements a simple logging package ... writes to standard err and prints
the date and time of each logged messaged ... Fatal functions call os.Exit(1) after writing
the log message ... the Panic function calls panic after writing the log message.

Run Result:
panic: open no-file.txt: The system cannot find the file specified.

goroutine 1 [running]:
main.main()
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTra
ining/21_error-handling/02_err-not-nil/05_panic/main.go:13 +0x6a
exit status 2
*/
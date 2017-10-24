package main

import (
	"os"
	"log"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// 	fmt.Println("err happened:", err)
		//	log.Println("err happened:", err)
		log.Fatalln(err)
		// 	panic(err)
	}
}

/*
Package log implements a simple logging package ... writes to standard err and prints
the date and time of each logged messaged ... Fatal functions call os.Exit(1) after writing
the log message ... the Panic function calls panic after writing the log message.

Run Result:
2017/10/24 02:04:34 open no-file.txt: The system cannot find the file specified.
exit status 1
*/
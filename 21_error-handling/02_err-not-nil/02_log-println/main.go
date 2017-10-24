package main

import (
	"os"
	"log"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//	fmt.Println("err happened:", err)
		log.Println("err happened:", err)
		//	log.Fatalln(err)
		//	panic(err)
	}
}

/*
Package log implements a simple logging package ... writes to standard error
and prints the date and time of each logged message ... the Fatal functions
call os.Exit(1) after writing the log message ... the Panic functions call panic
writing the log message

log.Println calls Output to print to the standard logger. Arguments are handled in
the manner of fmt.Println.

Run Results:
2017/10/24 01:36:19 err happened: open no-file.txt: The system cannot find the file sp
ecified.
*/
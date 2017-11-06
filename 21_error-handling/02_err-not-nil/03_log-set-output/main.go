package main

import (
	"os"
	"log"
	"fmt"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

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
Run Result:
new log file is created. func init() run initializations or setup before running main()

*/
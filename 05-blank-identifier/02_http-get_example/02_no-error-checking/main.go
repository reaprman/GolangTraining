package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://mcleods.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}

/*
Run: Error not caught fatal error

panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x48 pc=0x5d5dc9]

goroutine 1 [running]:
main.main()
        C:/Users/ryang/Documents/Projects/goprojects/src/github.com/reaprman/GolangTra
ining/05-blank-identifier/02_http-get_example/02_no-error-checking/main.go:11 +0x59
exit status 2
*/

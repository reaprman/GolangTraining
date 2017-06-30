package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.mcleods.com/")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err == nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}

/*
Run: Error caught

2017/06/30 01:16:31 Get http://www.mcleods.com/: dial tcp: lookup www.mcleods.com: get
addrinfow: This is usually a temporary error during hostname resolution and means that
 the local server did not receive a response from an authoritative server.
exit status 1

*/

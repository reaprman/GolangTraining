package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

func main() {
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}

/*
The tag json:"wisdom score", will turn the value of Age into
json variable wisdom score as seen below.

The "-" json option will not include this variable when
marhsal is run on p1.

Run Result:
{"First":"James","wisdom score":20}
*/

package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	fido := dog{animal{"BARK!!!"}, false}
	mipsy := cat{animal{"Hiss!!"}, true}
	shadow := dog{animal{"woof"}, true}
	critters := []interface{}{fido, mipsy, shadow}
	fmt.Printf("%T: %v", critters,critters)
}

/*
Run Result:
[]interface {}: [{{BARK!!!} false} {{Hiss!!} true} {{woof} true}]
*/

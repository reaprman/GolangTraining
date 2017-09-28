package main

import "fmt"

type animals interface{}

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

func specs(a interface{}) {
	fmt.Printf("%T: %v\n", a, a)
}

func main() {
	fido := dog{animal{"woof"}, true}
	mipsy := cat{animal{"meow"}, true}
	specs(fido)
	specs(mipsy)
}

/*
Empty interfaces allow you to pass anything into them

Run Result:
main.dog: {{woof} true}
main.cat: {{meow} true}
*/

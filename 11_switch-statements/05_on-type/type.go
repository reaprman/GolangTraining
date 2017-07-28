package 

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x,(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case Contact:
			fmt.Println("contact")
		default:
			fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("McLeod")
	var t = Contact{"Godd to see you,",}
}
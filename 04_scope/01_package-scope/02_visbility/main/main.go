package main

import (
	"fmt"
	"github.com/reaprman/GolangTraining/04_scope/01_package-scope/02_visbility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}

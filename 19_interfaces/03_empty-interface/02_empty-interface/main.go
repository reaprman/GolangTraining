package main

import "fmt"

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    int
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func (v vehicle) Specs() {
	fmt.Printf("Seats %v, max speed %v, color %v\n", v.Seats, v.MaxSpeed, v.Color)
}

func main() {
	prius := car{}
	tacoma := car{}
	bmw528 := car{}
	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}
	sanger := boat{}
	nautique := boat{}
	malibu := boat{}
	rides := []vehicles{prius, tacoma, bmw528, boeing747, boeing757, boeing767, sanger, nautique, malibu}

	for key, value := range rides {
		fmt.Printf("%T: %v - %v\n", rides[key], key, value)
	}
}
/*
Run Result:
main.car: 0 - {{0 0 0} 0 0}
main.car: 1 - {{0 0 0} 0 0}
main.car: 2 - {{0 0 0} 0 0}
main.plane: 3 - {{0 0 0} false}
main.plane: 4 - {{0 0 0} false}
main.plane: 5 - {{0 0 0} false}
main.boat: 6 - {{0 0 0} 0}
main.boat: 7 - {{0 0 0} 0}
main.boat: 8 - {{0 0 0} 0}
*/
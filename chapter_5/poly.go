package main

import "fmt"

type VehicleOne interface {
	move()
}

type Bibika struct {
	model string
}

type Airplane struct {
	model string
}

func (b *Bibika) move() {
	fmt.Println(b.model, "moves")
}

func (a *Airplane) move() {
	fmt.Println(a.model, "flies")
}

func Poly() {
	bibika := Bibika{"Tesla"}
	airplane := Airplane{"Boeing"}

	vehicles := [...]VehicleOne{&bibika, &airplane}
	for _, vehicle := range vehicles {
		vehicle.move()
	}
}

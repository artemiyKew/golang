package main

import "fmt"

type mile int
type kilometr int

type library []string

func DerivedTypes() {
	fmt.Println("Derived types!")
	var distance mile = 5
	fmt.Println(distance)
	distance += 5
	fmt.Println(distance)

	distanceToEnemy(distance)

	// var dist2 kilometr = 7
	// distanceToEnemy(dist2)

	var myLib library = library{"Book1", "Book2", "Book3"}
	printBooks(myLib)
}

func printBooks(lib library) {
	for _, value := range lib {
		fmt.Println(value)
	}
}

func distanceToEnemy(dist mile) {
	fmt.Println("Distance to enemy")
	fmt.Println(dist, "miles")
}

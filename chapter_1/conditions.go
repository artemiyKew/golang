package main

import "fmt"

func Conditions() {
	fmt.Println("\nУсловные конструкции!")
	var (
		a int = 5
		b int = 4
	)
	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a == b")
	}

	val := 6
	switch val {
	case 3:
		fmt.Println("val = 3")
	case 4:
		fmt.Println("val = 4")
	case 5:
		fmt.Println("val = 5")
	default:
		fmt.Println("I don`t know what is it!)")
	}
}

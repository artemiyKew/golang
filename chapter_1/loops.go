package main

import "fmt"

func Loops() {
	fmt.Println("\nЦиклы!")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	users := [2]string{"vasya", "petya"}

	for idx, val := range users {
		fmt.Println(idx, val)
	}
	for _, val := range users {
		fmt.Println(val)
	}
}

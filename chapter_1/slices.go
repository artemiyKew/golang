package main

import "fmt"

func Slices() {
	var users []string
	users = []string{"Vasya", "Petya"}

	users2 := make([]string, 0, 5)
	users2 = append(users2, "Bob")
	users2 = append(users2, users...)
	users2 = append(users2, users...)
	users2 = append(users2, users...)
	for index, value := range users2 {
		fmt.Println(index, ":", value)
	}
	fmt.Println(users2)
}

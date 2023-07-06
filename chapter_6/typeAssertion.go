package main

import "fmt"

func TypeAssertion() {
	var i interface{} = "Hello"
	c, ok := i.(string)
	fmt.Println(c, ok)

	var value interface{} = 2023

	a, test := value.(string)
	if test {
		fmt.Println("String value Found")
		fmt.Println(a)
	} else {
		fmt.Println("String value not found")
	}
}

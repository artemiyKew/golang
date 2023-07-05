package main

import (
	"chapter_4/computation"
	"fmt"
	"math"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(math.Sqrt(16))
	// fmt.Println(Factorial(10))
	fmt.Println(computation.Factorial(5))
	quote.Hello()
}

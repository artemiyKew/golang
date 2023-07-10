package main

import (
	"fmt"
	"runtime"
	"time"
)

func Goroutines() {
	fmt.Println("Hello")
	// runtime.GOMAXPROCS(1) // максимальное количество горутин
	for i := 1; i < 7; i++ {
		go func(n int, str string) {
			result := 1
			for j := 1; j <= n; j++ {
				result *= j
			}
			fmt.Println(n, "-", result)
		}(i, "abs")
	}

	fmt.Println("The end")
	fmt.Println(runtime.NumCPU())

	go showNumbers(10)

	time.Sleep(time.Second) //! заставляет планировщик запустить другие горутины если они есть
	// runtime.Gosched() // ! заставляет планировщик запустить другие горутины если они есть

	fmt.Println("Exit")
}

func factorial(n int) {
	if n < 1 {
		fmt.Println("Invalid input number")
		return
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)
}

func showNumbers(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	_ = wg
	go func() {

		fmt.Println("hello form anonymous func")
		// time.Sleep(time.Second)
		// fmt.Println("hello 2")
	}()

	go printHello()
	var p printer
	go p.printHello()
}

func printHello() {
	fmt.Println("hello from printHello")
}

type printer struct{}

func (p printer) printHello() {
	fmt.Println("hello from struct method")
}

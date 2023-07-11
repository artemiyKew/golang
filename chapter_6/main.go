package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello")
	// TypeAssertion()

	//Goroutines()
	// Sync()
	// Channel()
	Select()
	// counter := Counter{}

	// counter.Increment()
	// fmt.Println(counter.GetValue())
}

type Counter struct {
	mutex sync.RWMutex
	value int
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.value
}

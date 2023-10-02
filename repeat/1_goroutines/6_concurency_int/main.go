package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int
	mu := &sync.Mutex{}
	for i := 0; i < 1e5; i++ {
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second / 2)
	fmt.Println(count)
}

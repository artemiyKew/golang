package main

import (
	"sync"
	"time"
)

func main() {
	m := make(map[int]int)

	mu := &sync.Mutex{}

	go func() {
		for i := 0; i < 10000; i++ {
			mu.Lock()
			m[0] = i
			mu.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			mu.Lock()
			if m[0] < 0 {
			}
			mu.Unlock()
		}
	}()

	time.Sleep(time.Second / 2)
}

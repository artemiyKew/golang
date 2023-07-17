package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func Atomic() {
	addMutex()
	addAtomic()
	storeLoadSwap()
}

func addMutex() {
	start := time.Now()

	var (
		counter int64
		wg      = &sync.WaitGroup{}
		mu      = &sync.Mutex{}
	)
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	log.Println(time.Since(start).Seconds(), counter)
}

func addAtomic() {
	start := time.Now()

	var (
		counter int64
		wg      = &sync.WaitGroup{}
	)

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	log.Println(time.Since(start).Seconds(), counter)
}

func storeLoadSwap() {
	var counter int64

	fmt.Println(atomic.LoadInt64(&counter))

	atomic.StoreInt64(&counter, 5)
	fmt.Println(atomic.LoadInt64(&counter))

	fmt.Println(atomic.SwapInt64(&counter, 10))
	fmt.Println(atomic.LoadInt64(&counter))
}

func compareAndSwap() {

}

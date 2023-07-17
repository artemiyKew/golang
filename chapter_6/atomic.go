package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func Atomic() {
	// addMutex()
	// addAtomic()
	// storeLoadSwap()
	// compareAndSwap()
	atomicVal()
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
			defer mu.Unlock()
			counter++
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
	var (
		counter int64
		wg      = &sync.WaitGroup{}
	)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			if !atomic.CompareAndSwapInt64(&counter, 0, 1) {
				return
			}
			log.Println("Swapped goroutine numbers is ", i)
		}(i)
	}
	wg.Wait()
	fmt.Println(counter)
}

func atomicVal() {
	var (
		value atomic.Value
	)

	value.Store(1)
	fmt.Println(value.Load())

	fmt.Println(value.Swap(5))
	fmt.Println(value.Load())

	fmt.Println(value.CompareAndSwap(5, 10))
	fmt.Println(value.Load())
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func Sync() {
	// withoutWait()
	// withWait()
	// wrongAdd()
	// writeWithoouConcurrent()
	// writeWithoutMutex()
	// writeWithMutex()
	readWithMutex()
	readWithRWMutex()
}

func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}
	fmt.Println("exit")
}

func withWait() {
	var wg sync.WaitGroup
	// wg.Add(10) можно сделать так, если нам известно кол-во выполняемых задач
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i + 1)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("exit")
}

func wrongAdd() {
	var wg sync.WaitGroup
	// wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			// wg.Add(1) // так нельзя, тк горутина может не выполнится, потому что она добавляет себя во время создания, а не до создания
			defer wg.Done()
			fmt.Println(i + 1)

		}(i)
	}

	wg.Wait()
	fmt.Println("exit")
}

func writeWithoouConcurrent() {
	start := time.Now()
	var counter int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond)
		counter++
	}
	fmt.Println(counter)
	fmt.Println(time.Since(start).Seconds())
}

func writeWithoutMutex() {
	start := time.Now()

	var counter int
	var wg sync.WaitGroup

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		// wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			counter++
		}()
	}
	// идет потеря из-за DataRace -> readme.md
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Since(start).Seconds())
}

func writeWithMutex() {
	start := time.Now()

	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		// wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Since(start).Seconds())
}

func readWithMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			time.Sleep(time.Nanosecond)
			_ = counter
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()
			mu.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Since(start).Seconds())
}

func readWithRWMutex() {
	start := time.Now()

	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.RWMutex
	)

	wg.Add(100)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			mu.RLock()
			time.Sleep(time.Nanosecond)
			_ = counter
			mu.RUnlock()
		}()

		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Since(start).Seconds())
}

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func ErrGroup() {
	// chanAsPromise()
	// chanAsMutex()
	// withoutErrGroup()
	errGroup()
}

func makeRequest(i int) <-chan string {
	responseChan := make(chan string)

	go func() {
		time.Sleep(time.Second)
		responseChan <- fmt.Sprintf("Response number %d", i)
	}()

	return responseChan
}

func chanAsPromise() {
	firstResponseChan := makeRequest(1)
	secondResponseChan := makeRequest(2)
	//do smth else
	fmt.Println("non blocking")

	fmt.Println(<-firstResponseChan, <-secondResponseChan)
}

func chanAsMutex() {
	var counter int
	mutexChan := make(chan struct{}, 1)
	wg := &sync.WaitGroup{}
	// mu := &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mutexChan <- struct{}{}

			// mu.Lock() тоже самое что ко выше
			counter++
			// mu.Unlock()
			<-mutexChan
		}()
	}

	wg.Wait()
	fmt.Println(counter)

}

func withoutErrGroup() {
	var err error
	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(4)

	go func() {
		time.Sleep(time.Second)
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("first started")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("second started")
			err = fmt.Errorf("some err")
			cancel()
		}
	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("third started")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		time.Sleep(time.Millisecond)
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("fourth started")
		}
	}()

	wg.Wait()
	fmt.Println(err)
}

func errGroup() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			return nil
		default:
			fmt.Println("first started")
			time.Sleep(time.Second)
			return nil
		}
	})

	g.Go(func() error {
		fmt.Println("second started")
		return fmt.Errorf("unexpected error in request 2")
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("third started")
			time.Sleep(time.Second)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

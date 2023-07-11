package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Select() {
	// baseSelect()
	greacefulShutDown()
}

func baseSelect() {
	bufferedChannel := make(chan string, 3)
	bufferedChannel <- "first"

	select {
	case str := <-bufferedChannel:
		fmt.Println("read", str)
	case bufferedChannel <- "second":
		fmt.Println("write", <-bufferedChannel, <-bufferedChannel)
	}

	unbufChan := make(chan int)

	go func() {
		time.Sleep(time.Second)
		unbufChan <- 1
	}()

	select {
	// case bufferedChannel <- "third":
	// 	fmt.Println("unblocking writing")
	case val := <-unbufChan:
		fmt.Println("blocking writing", val)
	case <-time.After(time.Millisecond * 500):
		fmt.Println("time`s up")
		// default:
		// 	fmt.Println("default case")
	}

	resutlChan := make(chan int)
	timer := time.After(time.Second)

	go func() {
		defer close(resutlChan)

		for i := 0; i < 1000; i++ {
			select {
			case <-timer:
				fmt.Println("time`s up")
				return
			default:
				time.Sleep(time.Nanosecond)
				resutlChan <- i
			}
		}
	}()

	for v := range resutlChan {
		fmt.Println(v)
	}
}

func greacefulShutDown() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(10 * time.Second)

	select {
	case <-timer:
		fmt.Println("time`s up")
		return
	case sig := <-sigChan:
		fmt.Println("Stopped by signal", sig)
		return
	}

}

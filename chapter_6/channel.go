package main

import (
	"fmt"
	"sync"
	"time"
)

func Channel() {
	// nilChannel()
	// unbufferedChannel()
	// bufferedChannel()
	forRange()
}

func nilChannel() {
	var nilChannel chan int
	fmt.Printf("Len: %d Cap: %d\n", len(nilChannel), cap(nilChannel))

	// *write to nil channel blocks forever
	// nilChannel <- 1

	// *read from nil channel blocks forever
	// <-nilChannel

	// * closing nil channel will raise a panic
	// close(nilChannel)
}

func unbufferedChannel() {
	unbufferedChannel := make(chan int)
	// unbufferedChannel := make(chan<- int) //!Только запись в канал(канал для записи)
	// unbufferedChannel := make(<-chan int) //!Только чтение из канала(канал для чтения )
	fmt.Printf("Len: %d Cap: %d\n", len(unbufferedChannel), cap(unbufferedChannel))

	// * blocks until smb reads
	// unbufferedChannel <- 1

	// * blocks until smb write
	// <-unbufferedChannel

	// * blocks reading then write
	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)
		// <-chanForWriting
		unbufferedChannel <- 3
	}(unbufferedChannel)

	value := <-unbufferedChannel
	fmt.Println(value)

	// * blocks writing then read
	go func(chanForReading <-chan int) {
		time.Sleep(time.Second)
		value := <-chanForReading
		fmt.Println(value)
	}(unbufferedChannel)
	unbufferedChannel <- 2

	// close(unbufferedChannel)
	// * close channel

	// go func() {
	// 	time.Sleep(time.Millisecond * 500)
	// 	close(unbufferedChannel)
	// }()
	// go func() {
	// 	time.Sleep(time.Second)
	// 	fmt.Println(<-unbufferedChannel)
	// }()
	// !panic
	// unbufferedChannel <- 4
	close(unbufferedChannel)
}

func bufferedChannel() {
	bufferedChannel := make(chan int, 2)
	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))
	bufferedChannel <- 2
	bufferedChannel <- 3
	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))

	// * blocks to write, buffer is full
	// bufferedChannel <- 4

	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))
	bufferedChannel <- 4
	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))
	fmt.Println(<-bufferedChannel)

	close(bufferedChannel)
}

func forRange() {
	bufferedChannel := make(chan int, 4)
	numbers := []int{1, 2, 3, 4, 5}

	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
			fmt.Println("Отправлено:", num)
		}
		close(bufferedChannel)
	}()

	for {
		v, ok := <-bufferedChannel
		fmt.Println("Получено:", v)
		fmt.Println(v, ok)
		if !ok {
			break
		}
	}

	fmt.Println("\nbufferedChannel = make(chan int, 3)")
	bufferedChannel = make(chan int, 3)

	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
			fmt.Println("Отправлено:", num)
		}
		close(bufferedChannel)
	}()

	for v := range bufferedChannel {
		fmt.Println("buffered", v)
	}

	fmt.Println("\nunbufferedChannel := make(chan int)")
	unbufferedChannel := make(chan int)

	var wg sync.WaitGroup
	go func() {
		for _, num := range numbers {
			wg.Add(1)
			fmt.Println("Отправлено:", num)
			unbufferedChannel <- num
		}
		close(unbufferedChannel)
	}()

	for v := range unbufferedChannel {
		fmt.Println("unbuffered", v)
	}
	wg.Wait()
}

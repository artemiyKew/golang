package search

import (
	"fmt"
	"runtime"
	"time"
)

func Search() {
	radius := 5
	fmt.Println(runtime.NumCPU())
	driver := make(chan Driver)

	go func() {
		defer close(driver) // Закрыть канал driver при завершении goroutine
		for {
			timer := time.After(time.Second * 2)
			select {
			case <-timer:
				radius *= 2
				if radius >= 10 {
					d := Driver{name: "Victor", phone: "+79297987666"}
					driver <- d
					return // Завершить goroutine после отправки данных
				}
			}
		}
	}()
	value := <-driver
	fmt.Println(value)
}

type Driver struct {
	name  string
	phone string
}

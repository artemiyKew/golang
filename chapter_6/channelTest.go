package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func ChannelTest() {
	var (
		wg       = &sync.WaitGroup{}
		services = []string{"Yandex Go", "CityMobil", "YO", "Volga"}
		resultCh = make(chan string, len(services))
		// winnerCh    = make(chan string)
		ctx, cancel = context.WithCancel(context.Background())
		winner      string
		ok          bool
	)

	defer cancel()
	for i := range services {
		sv := services[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			searchService(ctx, sv, resultCh)
		}()
	}

	go func() {
		winner, ok = <-resultCh
		fmt.Println(ok)
		cancel()
	}()
	wg.Wait()
	if ok {
		log.Printf("found car in %q", winner)
	}
	close(resultCh)
}

func searchService(ctx context.Context, service string, resultCh chan<- string) {
	time.Sleep(time.Second * 1)
	for {
		select {
		case <-ctx.Done():
			log.Printf("stopped the search in %q (%v)", service, ctx.Err())
			return
		default:
			r := rand.Float64()
			fmt.Println(r, ":", service)
			if r > 0.75 {
				resultCh <- service
				return
			}
			continue
		}
	}
}

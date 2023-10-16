package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go handelSignals(cancel)

	if err := startServer(ctx); err != nil {
		log.Fatal(err)
	}
}

func handelSignals(cancel context.CancelFunc) {
	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		fmt.Println("Stopped by signal", sig)
		cancel()
	}
}

func startServer(ctx context.Context) error {
	log.Println("server started")
	laadr, err := net.ResolveTCPAddr("tcp", ":8080")

	if err != nil {
		return nil
	}

	l, err := net.ListenTCP("tcp", laadr)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		select {
		case <-ctx.Done():
			log.Println("server stopped")
			return nil
		default:
			if err := l.SetDeadline(time.Now().Add(time.Second)); err != nil {
				return err
			}
			_, err := l.Accept()
			if err != nil {
				if os.IsTimeout(err) {
					continue
				}
				return err
			}

			log.Println("new client connected")
		}
	}
}

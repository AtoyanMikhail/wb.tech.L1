package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("1. Stop by condition:")
	done := false
	go func() {
		for {
			if done {
				fmt.Println("Goroutine stopped by condition")
				return
			}
			fmt.Println("Working...")
			time.Sleep(300 * time.Millisecond)
		}
	}()
	time.Sleep(time.Second)
	done = true
	time.Sleep(400 * time.Millisecond)
	fmt.Println()

	fmt.Println("2. Stop via channel:")
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Goroutine stopped via channel")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	time.Sleep(time.Second)
	close(stop)
	time.Sleep(400 * time.Millisecond)
	fmt.Println()

	fmt.Println("3. Stop via context.Context:")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stopped via context")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	time.Sleep(time.Second)
	cancel()
	time.Sleep(400 * time.Millisecond)
	fmt.Println()

	fmt.Println("4. Stop via runtime.Goexit():")
	finished := make(chan struct{})
	go func() {
		defer func() { close(finished) }()
		fmt.Println("Working...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Calling runtime.Goexit()")
		runtime.Goexit()
		// This code will not be executed
		fmt.Println("This text will not be printed")
	}()
	<-finished
	fmt.Println("Goroutine stopped via Goexit")
	fmt.Println()
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Системное прерывание разблокирует горутину 1 и отменит контекст, 
// после чего оба воркера прослушают отмену контекста в select и завершатся
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	// Горутина 1
	go func(){
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt)
		<-exit
		cancel()
	}()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Элегантное завершение 1")
				return
			case <-time.After(1*time.Second):
				fmt.Println("Бизнес логика 1")
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Элегантное завершение 2")
				return
			case <-time.After(1*time.Second):
				fmt.Println("Бизнес логика 2")
			}
		}
	}()

	wg.Wait()
}

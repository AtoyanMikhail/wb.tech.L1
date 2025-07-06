package main

import (
	"fmt"
	"sync"
)

func main() {
	var workers int
	fmt.Scanf("%d", &workers)

	wg := sync.WaitGroup{}
	wg.Add(workers)
	ch := make(chan int)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			data := <-ch
			fmt.Println(data)
		}()
	}

	// Если не ограничить кол-во данных, то main заблокируется
	// попыткой записи в переполненный канал
	for i := 0; i < workers; i++ {
		ch <- i
	}
	wg.Wait()
}

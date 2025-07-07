package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	numGoroutines = 10
	incrementsPerGoroutine = 1000
)

type Counter struct {
	value int64
}

func (c *Counter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	wg := &sync.WaitGroup{}
	counter := &Counter{}

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Value())
}

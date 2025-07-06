package main

import "sync"

type concurrentMap struct {
	mu *sync.RWMutex //
	m map[int]int
}

func (cm *concurrentMap) Set(key, value int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.m[key] = value
}

func (cm *concurrentMap) Get(key int) (value int, ok bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	value, ok = cm.m[key] 
	return
}

func main() {
	m := concurrentMap{
		&sync.RWMutex{},
		map[int]int{},
	}

	wg := &sync.WaitGroup{}
	wg.Add(1000)

	for i := range 1000 {
		go func(i int) {
			defer wg.Done()
			m.Set(0, i)
		}(i)
	}

	wg.Wait()
}

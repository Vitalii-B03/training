package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
}

func (c *Counter) Increment() int {
	c.value++
	return c.value
}

func concurrentSafeCounter() int {
	counter := Counter{}
	var mutex sync.Mutex
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			counter.Increment()
			mutex.Unlock()
		}()
	}
	wg.Wait()
	return counter.value
}
func main() {
	fmt.Println(concurrentSafeCounter())
}

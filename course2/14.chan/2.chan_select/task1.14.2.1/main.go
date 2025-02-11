package main

import (
	"fmt"
	"sync"
	"time"
)

func trySend(ch chan int, v int) bool {
	time.Sleep(1 * time.Second)
	select {
	case ch <- v:
		return true
	default:
		return false
	}
}

func main() {
	ch := make(chan int, 1)
	v := 8
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(trySend(ch, v))
	}()
	wg.Wait()
	close(ch)
}

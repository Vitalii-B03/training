package main

import (
	"fmt"
	"sync"
)

func waitGroupExample(goroutines ...func() string) string {
	var wg sync.WaitGroup
	var result string
	var resChan = make(chan string, len(goroutines))
	for i, goroutine := range goroutines {
		wg.Add(1)
		go func(i int, goroutine func() string) {
			defer wg.Done()
			resChan <- fmt.Sprintf("goroutine %d done", i)
		}(i, goroutine)
	}
	wg.Wait()
	close(resChan)
	for res := range resChan {
		result += res + "\n"
	}
	return result
}

func main() {
	count := 1000
	goroutines := make([]func() string, count)

	for i := 0; i < count; i++ {
		j := i
		goroutines[i] = func() string {
			return fmt.Sprintf("goroutine %d", j)
		}
	}
	fmt.Println(waitGroupExample(goroutines...))
}

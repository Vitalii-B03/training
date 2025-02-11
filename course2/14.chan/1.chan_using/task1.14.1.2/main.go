package main

import (
	"fmt"
	"time"
)

func generateData(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()
	return out
}
func main() {
	data := generateData(10)
	go func() {
		time.Sleep(1 * time.Second)
		close(data)
	}()
	for num := range data {
		fmt.Println(num)
	}
}

package main

import (
	"fmt"
	"time"
)

func timeout(timeout time.Duration) func() bool {
	done := make(chan bool, 1)
	go func() {
		time.Sleep(timeout)
		done <- true
	}()
	return func() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	}
}
func main() {
	timeoutFunc := timeout(3 * time.Second)
	since := time.NewTimer(3050 * time.Millisecond)
	for {
		select {
		case <-since.C:
			fmt.Println("Функция не выполнена вовремя")
			return
		default:
			if timeoutFunc() {
				fmt.Println("Функция выполнена вовремя")
				return
			}
		}
	}
}

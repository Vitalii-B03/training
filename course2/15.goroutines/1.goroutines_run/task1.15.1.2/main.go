package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем новый тикер с интервалом 1 секунда
	ticker := time.NewTicker(1 * time.Second)

	data := NotifyEvery(ticker, 5*time.Second, "Таймер сработал")

	for v := range data {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")
}

func NotifyEvery(ticker *time.Ticker, d time.Duration, message string) <-chan string {
	var result = make(chan string)
	timer := time.NewTimer(d)
	go func() {
		for {
			select {
			case <-ticker.C:
				result <- message
			case <-timer.C:
				close(result)
			}
		}
	}()
	return result
}

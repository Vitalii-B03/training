package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

// Пример структуры счетчика
type Counter struct {
	count int64
}

// Функция для увеличения значения счетчика на 1
func (c *Counter) Increment() {
	mutex.Lock()
	c.count++
	mutex.Unlock()
}

// Функция для получения текущего значения счетчика
func (c *Counter) GetCount() int64 {
	// Ваш код для получения текущего значения счетчика
	return c.count
}
func main() {
	counter := &Counter{}
	for i := 0; i < 10; i++ {
		counter.Increment()
	}
	fmt.Println(counter.GetCount())
}

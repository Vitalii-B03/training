package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)
	myPanic(ch)
	fmt.Println(<-ch)

}

func myPanic(ch chan string) {
	defer func() {
		if err := recover(); err != nil {
			ch <- fmt.Sprint(err)
		}
	}()
	panic("my panic message")

}

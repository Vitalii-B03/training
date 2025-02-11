package main

import "fmt"

func adder(initial int) func(int) int {
	a := initial
	return func(b int) int {
		return a + b
	}
}

func main() {
	// пример использования функции adder
	addTwo := adder(2)
	result := addTwo(3)
	fmt.Println(result) // выводит 5
}

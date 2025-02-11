package main

import "fmt"

func calculate(a int, b int) (int, int, int, int, int) {
	var sum int
	var difference int
	var product int
	var quotient int
	var remainder int
	sum = a + b
	difference = a - b
	product = a * b
	quotient = a / b
	remainder = a % b
	return sum, difference, product, quotient, remainder
}

func main() {
	sum, difference, product, quotient, remainder := calculate(10, 3)
	fmt.Printf("sum = %d difference = %d product = %d quotient = %d remainder = %d\n", sum, difference, product, quotient, remainder)
}

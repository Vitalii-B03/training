package main

import "fmt"

func bitwiseAnd(a, b int) int {
	c := a & b
	return c
}

func bitwiseOr(a, b int) int {
	c := a | b
	return c
}

func bitwiseXor(a, b int) int {
	c := a ^ b
	return c
}

func bitwiseLeftShift(a, b int) int {
	c := a << b
	return c
}

func bitwiseRightShift(a, b int) int {
	c := a >> b
	return c
}

func main() {
	fmt.Printf("a & b = %d\n", bitwiseAnd(1, 7))
	fmt.Printf("a | b = %d\n", bitwiseOr(1, 1))
	fmt.Printf("a ^ b = %d\n", bitwiseXor(1, 1))
	fmt.Printf("a << b = %d\n", bitwiseLeftShift(1, 1))
	fmt.Printf("a >> b = %d\n", bitwiseRightShift(1, 1))
}

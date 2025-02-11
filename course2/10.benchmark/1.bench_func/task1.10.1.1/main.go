package main

func Fibonacci(n int) int {
	if n < 2 {
		return n
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}

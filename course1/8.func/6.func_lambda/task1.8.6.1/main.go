package main

import "fmt"

func Sum(a ...int) int {
	sum := 0
	for _, n := range a {
		sum += n

	}

	return sum
}

func Mul(a ...int) int {
	mul := 1
	for _, n := range a {
		mul *= n
	}
	return mul
}

func Sub(a ...int) int {
	sub := a[0]
	for _, n := range a[1:] {
		sub -= n
	}
	return sub
}

func MathOperate(op func(a ...int) int, a ...int) int {
	return op(a...)
}
func main() {
	fmt.Println(MathOperate(Sum, 1, 1, 3))
	fmt.Println(MathOperate(Mul, 1, 7, 3))
	fmt.Println(MathOperate(Sub, 13, 2, 3))
}

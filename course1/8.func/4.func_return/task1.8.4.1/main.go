package main

import "fmt"

func DivideAndRemainder(a, b int) (int, int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("check zero argument")
		}
	}()
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
func main() {
	fmt.Println(DivideAndRemainder(4, 2))
}

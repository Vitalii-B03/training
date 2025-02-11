package main

import (
	"math/big"
)

func main() {
	a, b, number := 1, 3, 12
	numbers := []int{6, 2, 3, 4, 5}
	strs := []string{"task", "1.5.1.1"}
	ConcatenateStrings(strs)
	Add(a, b)
	Max(numbers)
	IsPrime(number)
}
func Add(a, b int) *int {
	c := a + b
	return &c
}
func Max(numbers []int) *int {
	max := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return &max
}
func IsPrime(number int) *bool {
	if big.NewInt(int64(number)).ProbablyPrime(0) {
		num := true
		return &num
	} else {
		num1 := false
		return &num1
	}
}
func ConcatenateStrings(strs []string) *string {
	var strs1 string
	for i := 0; i < len(strs); i++ {
		strs1 += strs[i]
	}
	return &strs1
}

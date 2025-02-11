package main

import "fmt"

func main() {
	fctrl := "qwerewq"
	fmt.Println(isPalindrome(&fctrl))
	a := 6
	fmt.Println(Factorial(&a))
	numb := []int{1, 3, 5, 4, 4, 5, 7}
	targ := 4
	fmt.Println(CountOccurrences(&numb, &targ))
	str := "Revers"
	fmt.Println(ReverseString(&str))
}
func Factorial(n *int) int {
	b := *n
	for i := *n - 1; i > 0; i-- {
		b = b * i
	}
	return b
}
func isPalindrome(str *string) bool {
	runes := []rune(*str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	if string(runes) == *str {
		return true
	} else {
		return false
	}
}
func CountOccurrences(numbers *[]int, target *int) int {
	j := 0
	for i := 0; i < len(*numbers); i++ {
		if (*numbers)[i] == *target {
			j = j + 1
		}
	}
	return j
}

func ReverseString(str *string) string {
	runes := []rune(*str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

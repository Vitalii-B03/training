package main

import "fmt"

func main() {
	a := 3
	var a1 *int = &a
	str := "Hello"
	mutate(a1)
	ReverseString(&str)
	fmt.Println(str)
}
func mutate(a1 *int) {
	*a1 = 42
}
func ReverseString(str *string) {
	runes := []rune(*str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*str = string(runes)
}

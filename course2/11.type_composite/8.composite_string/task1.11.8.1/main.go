package main

import (
	"fmt"
	"unicode/utf8"
)

func countBytes(s string) int {

	return len(s)
}
func countSymbols(s string) int {

	return utf8.RuneCountInString(s)
}
func main() {
	s := "Привет, мир!"
	fmt.Println(countBytes(s))
	fmt.Println(countSymbols(s))
}

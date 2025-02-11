package main

import (
	"fmt"
)

func ReverseString(s string) string {
	var rev string
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		rev += string(runes[i])
	}
	return rev
}
func main() {
	fmt.Println(ReverseString("xol"))
}

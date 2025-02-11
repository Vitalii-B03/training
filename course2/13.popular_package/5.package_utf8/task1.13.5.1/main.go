package main

import (
	"fmt"
	"unicode/utf8"
)

func countUniqueUTF8Chars(s string) int {
	count := make(map[string]struct{})
	var str string
	for _, ch := range s {
		count[string(ch)] = struct{}{}
	}
	for ch := range count {
		str += ch
	}
	return utf8.RuneCountInString(str)
}
func main() {
	fmt.Println(countUniqueUTF8Chars("Hello, 世界!"))
}

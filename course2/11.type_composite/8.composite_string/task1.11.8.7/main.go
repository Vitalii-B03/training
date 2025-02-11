package main

import "fmt"

func ReplaceSymbols(s string, old rune, new rune) string {
	str := []rune(s)
	oldB := old
	newB := new
	var result string
	for _, symbol := range str {
		if symbol == oldB {
			result += string(newB)
		} else {
			result += string(symbol)
		}
	}
	return result
}
func main() {
	result := ReplaceSymbols("Hello, world!", 'o', '0')
	fmt.Println(result)
}

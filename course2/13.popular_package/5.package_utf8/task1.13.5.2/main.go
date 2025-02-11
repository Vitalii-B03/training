package main

import (
	"fmt"
	"unicode"
)

func countRussianLetters(s string) map[rune]int {
	counts := make(map[rune]int)

	for _, char := range s {
		if isRussianLetter(char) {
			char = unicode.ToLower(char)
			counts[char]++
		}

	}
	return counts
}
func isRussianLetter(char rune) bool {
	return char >= 'а' && char <= 'я' || char >= 'А' && char <= 'Я' || char == 'ё' || char == 'Ё'
}
func main() {
	result := countRussianLetters("Привет, мир!")
	for key, value := range result {
		fmt.Printf("%c: %d ", key, value)
	}
}

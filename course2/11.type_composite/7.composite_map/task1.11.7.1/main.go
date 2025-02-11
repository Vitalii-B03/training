package main

import (
	"fmt"
	"strings"
)

func countWordOccurrences(text string) map[string]int {
	var m = make(map[string]int)
	var t []string
	t = strings.Split(text, " ")
	for _, word := range t {
		m[word]++
	}
	return m
}
func main() {
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	occurrences := countWordOccurrences(text)

	for word, count := range occurrences {
		fmt.Printf("%s: %d\n", word, count)
	}
}

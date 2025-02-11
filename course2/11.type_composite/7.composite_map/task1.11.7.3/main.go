package main

import (
	"fmt"
	"strings"
)

func createUniqueText(text string) string {
	mapText := make(map[string]bool)
	var text1 []string
	var str string
	text1 = strings.Split(text, " ")
	for _, word := range text1 {
		mapText[word] = true
	}
	for _, t := range text1 {
		for key, _ := range mapText {
			if t == key {
				str += key + " "
				delete(mapText, key)
				fmt.Println(mapText)
			}
		}

	}
	str = str[:len(str)-1]
	return str
}
func main() {
	fmt.Println(createUniqueText("bar bar bar foo foo baz"))
}

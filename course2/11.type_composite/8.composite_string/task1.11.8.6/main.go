package main

import "fmt"

func CountVowels(s string) int {
	v := "АаУуОоИиЭэЫыЯяЮюЕеЁёAaEeIiOoUuYy"
	runeV := []rune(v)
	runeS := []rune(s)
	i := 0
	for _, V := range runeV {
		for _, S := range runeS {
			if V == S {
				i++
			}
		}
	}
	return i
}
func main() {
	count := CountVowels("Hello, wold!")
	fmt.Println(count)
}

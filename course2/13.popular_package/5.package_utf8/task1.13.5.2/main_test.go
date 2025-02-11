package main

import (
	"strconv"
	"testing"
)

func TestCountRussianLetters(t *testing.T) {
	exp := "Привет, мир!"
	val := map[rune]int{'в': 1, 'е': 1, 'т': 1, 'м': 1, 'п': 1, 'р': 2, 'и': 2}
	if len(countRussianLetters(exp)) != len(val) {
		t.Errorf("exp len:%d,res:%d", len(exp), len(val))
	}
	res := countRussianLetters(exp)
	for k, v := range countRussianLetters(exp) {
		if val[k] != v {
			t.Errorf("exp:%d,res:%d", res[k], val[k])
		}
	}
}
func TestIsRussianLetter(t *testing.T) {
	runeTxt := map[rune]bool{
		'в': true,
		'e': false,
		'т': true,
		'g': false,
		'п': true,
		'p': false,
		'и': true,
	}
	i := 0
	for k, v := range runeTxt {
		i++
		t.Run("test"+strconv.Itoa(i), func(t *testing.T) {
			if isRussianLetter(k) != v {
				t.Errorf("exp:%t,res:%t", v, isRussianLetter(k))
			}
		})
	}
}
func BenchmarkIsRussianLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countRussianLetters("Привет, мир!")
	}

}

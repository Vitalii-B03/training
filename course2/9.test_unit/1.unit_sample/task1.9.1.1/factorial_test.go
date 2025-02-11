package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	result := Factorial(0)
	if result != 1 {
		t.Errorf("Expecting %d, got %d", 1, result)
	}
	result2 := Factorial(1)
	if result2 != 1 {
		t.Errorf("Expecting %d, got %d", 1, result2)
	}
	result3 := Factorial(5)
	if result3 != 120 {
		t.Errorf("Expecting %d, got %d", 120, result3)
	}
}

package main

import "testing"

func BenchmarkFibonacci(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}
func TestFibonacci(t *testing.T) {
	tests := []struct {
		a, b int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
	}
	for _, tt := range tests {
		if result := Fibonacci(tt.a); result != tt.b {
			t.Errorf("Fibonacci(%d) = %d, want %d", tt.a, result, tt.b)
		}
	}
}

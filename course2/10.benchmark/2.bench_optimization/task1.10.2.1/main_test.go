package main

import "testing"

func BenchmarkFibonacciDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciDP(10)
	}
}
func TestFibonacciDP(t *testing.T) {
	a := []struct {
		n, r int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}
	for _, i := range a {
		result := FibonacciDP(i.n)
		result2 := FibonacciBinet(i.n)
		if result != i.r {
			t.Errorf("fibonacciDP(%d)=%d, want %d", i.n, result, i.r)
		}
		if result2 != i.r {
			t.Errorf("fibonacciBinet(%d)=%d, want %d", i.n, result, i.r)
		}
	}

}

func BenchmarkFibonacciBinet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciBinet(10)
	}
}

package main

import "testing"

func TestFibonacci(t *testing.T) {
	type testData struct {
		input    int
		expected int
	}
	testCases := []testData{
		{input: 5, expected: 5},
		{input: 6, expected: 8},
		{input: 7, expected: 13},
		{input: 8, expected: 21},
		{input: 9, expected: 34},
		{input: 10, expected: 55},
	}
	for _, tc := range testCases {
		result := Fibonacci(tc.input)
		if result != tc.expected {
			t.Errorf("Fibonacci(%d) вернул %d, а верный результат %d", tc.input, result, tc.expected)
		}
	}
}

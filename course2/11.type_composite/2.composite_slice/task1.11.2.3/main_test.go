package main

import "testing"

func TestBitwiseXOR(t *testing.T) {
	tests := []struct {
		slice  []int
		result int
	}{
		{[]int{1, 1}, 0},
		{[]int{0, 1}, 1},
	}
	for _, test := range tests {
		exp := bitwiseXOR(test.slice[0], test.slice[1])
		if exp != test.result {
			t.Errorf("bitwiseXOR(%v, %v) => %d, want %d", test.slice[0], test.slice[1], exp, test.result)
		}
	}
}
func TestFindSingleNumber(t *testing.T) {
	tests := []struct {
		slice  []int
		result int
	}{
		{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1}, 5},
		{[]int{6, 7, 8, 9, 10, 9, 8, 7, 6}, 10},
	}
	for _, test := range tests {
		exp := findSingleNumber(test.slice)
		if exp != test.result {
			t.Errorf("findSingleNumber(%v) => %d, want %d", test.slice, exp, test.result)
		}
	}
}

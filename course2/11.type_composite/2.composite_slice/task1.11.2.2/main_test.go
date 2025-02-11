package main

import "testing"

func TestMaxDifference(t *testing.T) {
	tests := []struct {
		slice []int
		want  int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
		{[]int{1, 2, 3}, 2},
		{[]int{6, 2, 3, 4}, 4},
	}
	for _, test := range tests {
		if MaxDifference(test.slice) != test.want {
			t.Errorf("MaxDifference(%v) = %d, want %d", test.slice, MaxDifference(test.slice), test.want)
		}
	}
}

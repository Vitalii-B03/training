package main

import "testing"

func TestSum(t *testing.T) {
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := 36
	result := sum(a)
	if result != b {
		t.Errorf("sum(%d) = %d; want %d", a, b, result)
	}
}
func TestAverage(t *testing.T) {
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := 4.5
	result := average(a)
	if result != b {
		t.Errorf("average(%d) = %f; want %f", a, b, result)
	}
}
func TestAverageFloat(t *testing.T) {
	a := [8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}
	b := 5.0
	result := averageFloat(a)
	if result != b {
		t.Errorf("averageFloat(%f) = %f; want %f", a, b, result)
	}
}

func TestReverse(t *testing.T) {
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := [8]int{8, 7, 6, 5, 4, 3, 2, 1}
	result := reverse(a)
	if result != b {
		t.Errorf("reverse(%d) = %d; want %d", a, b, result)
	}
}

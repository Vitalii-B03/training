package main

import "testing"

func TestSortDescInt(t *testing.T) {
	test := [8]int{5, 2, 8, 1, 9, 3, 7, 4}
	expected := [8]int{9, 8, 7, 5, 4, 3, 2, 1}
	result := sortDescInt(test)
	if result != expected {
		t.Errorf("sortDescInt returned %+v, expected %+v", result, expected)
	}
}
func TestSortAscInt(t *testing.T) {
	test := [8]int{5, 2, 8, 1, 9, 3, 7, 4}
	expected := [8]int{1, 2, 3, 4, 5, 7, 8, 9}
	result := sortAscInt(test)
	if result != expected {
		t.Errorf("sortAscInt returned %+v, expected %+v", result, expected)
	}
}
func TestSortDescFloat(t *testing.T) {
	test := [8]float64{5.5, 2.2, 8.8, 1.1, 9.9, 3.3, 7.7, 4.4}
	expected := [8]float64{9.9, 8.8, 7.7, 5.5, 4.4, 3.3, 2.2, 1.1}
	result := sortDescFloat(test)
	if result != expected {
		t.Errorf("sortDescFloat returned %+v, expected %+v", result, expected)
	}
}
func TestSortAscFloat(t *testing.T) {
	test := [8]float64{5.5, 2.2, 8.8, 1.1, 9.9, 3.3, 7.7, 4.4}
	expected := [8]float64{1.1, 2.2, 3.3, 4.4, 5.5, 7.7, 8.8, 9.9}
	result := sortAscFloat(test)
	if result != expected {
		t.Errorf("sortDescFloat returned %+v, expected %+v", result, expected)
	}
}

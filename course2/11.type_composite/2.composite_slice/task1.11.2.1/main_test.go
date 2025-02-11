package main

import (
	"reflect"
	"testing"
)

func TestGetSubSlice(t *testing.T) {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	st := 2
	en := 6
	expected := []int{3, 4, 5, 6}
	result := getSubSlice(test, st, en)
	if reflect.DeepEqual(result, expected) != true {
		t.Errorf("getSubSlice(%v, %v, %v) = %v, want %v", test, st, en, result, expected)
	}
}

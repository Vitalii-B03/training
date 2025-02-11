package main

import (
	"reflect"
	"testing"
)

func TestAppendInt(t *testing.T) {
	testSlice := []int{1, 2, 3}
	b := &testSlice
	tests := []struct {
		slice2  []int
		reSlice []int
	}{
		{[]int{4}, []int{1, 2, 3, 4}},
		{[]int{}, []int{1, 2, 3, 4}},
		{[]int{5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, test := range tests {
		t.Run("Добавление элементов", func(t *testing.T) {
			appendInt(b, test.slice2...)
			if reflect.DeepEqual(testSlice, test.reSlice) == false {
				t.Errorf("appendInt = %v; должен %v", testSlice, test.reSlice)
			}
		})

	}

}

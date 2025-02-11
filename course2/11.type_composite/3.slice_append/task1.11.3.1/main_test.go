package main

import (
	"reflect"
	"testing"
)

func TestAppendInt(t *testing.T) {
	tests := []struct {
		slices  []int
		slice2  []int
		reSlice []int
	}{
		{[]int{1, 2, 3}, []int{4}, []int{1, 2, 3, 4}},
		{[]int{}, []int{4}, []int{4}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
	}
	for _, test := range tests {
		t.Run("Добавление элементов", func(t *testing.T) {
			exp := appendInt(test.slices, test.slice2...)
			if reflect.DeepEqual(exp, test.reSlice) == false {
				t.Errorf("appendInt = %v; должен %v", exp, test.reSlice)
			}
		})

	}

}

package main

import (
	"reflect"
	"testing"
)

//	func generateSlice(size int) []float64 {
//		var s []float64
//		for i := 0; i < size; i++ {
//			s = append(s, float64(rand.Intn(size)))
//		}
//		return s
//	}
func TestAverage(t *testing.T) {
	expected := 2.0
	result := average(generateSlice(5))

	if reflect.DeepEqual(generateSlice(5), generateSlice(5)) != true {
		t.Errorf("generateSlice работает не правильно")
	}
	if result != expected {
		t.Errorf("функция average(5) дала результат %f, а долна %f", result, expected)
	}
}

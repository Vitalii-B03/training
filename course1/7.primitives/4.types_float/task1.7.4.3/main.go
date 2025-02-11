package main

import (
	"fmt"
	"math"
)

func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	round1 := math.Round(a*math.Pow(10, float64(decimalPlaces))) / math.Pow(10, float64(decimalPlaces))
	round2 := math.Round(b*math.Pow(10, float64(decimalPlaces))) / math.Pow(10, float64(decimalPlaces))
	if round1 == round2 {
		isEqual = true
	}
	difference = round1 - round2
	return isEqual, difference
}
func main() {
	fmt.Println(CompareRoundedValues(3.67866, 3.67566, 4))
}

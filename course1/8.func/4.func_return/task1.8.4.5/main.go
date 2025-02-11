package main

import (
	"fmt"
)

func CalculatePercentageChange(initialValue, finalValue float64) float64 {
	return ((finalValue - initialValue) / initialValue) * 100
}
func main() {
	fmt.Printf("%.2f%%\n", CalculatePercentageChange(75, 50))
}

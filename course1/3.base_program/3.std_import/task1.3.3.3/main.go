package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Sin(0.134), Cos(0.134))
}
func Sin(x float64) float64 {
	y := math.Sin(x)
	return y
}
func Cos(x float64) float64 {
	y := math.Cos(x)
	return y
}

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Floor(3.6574))
}
func Floor(x float64) float64 {
	y := math.Floor(x)
	return y
}

package main

import (
	"fmt"
	"math"
)

func hypotenuse(a, b float64) float64 {
	c := math.Sqrt(a*a + b*b)
	return c

}
func main() {
	fmt.Println(hypotenuse(6, 8))
}

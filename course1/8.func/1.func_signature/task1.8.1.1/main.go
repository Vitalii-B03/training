package main

import "fmt"

var CalculateCircleArea func(radius float64) float64 = func(radius float64) float64 {
	return 3.141592653589793 * radius * radius
}
var CalculateRectangleArea func(width, height float64) float64 = func(width, height float64) float64 {
	return height * width
}
var CalculateTriangleArea func(base, height float64) float64 = func(base, height float64) float64 {
	return (base * height) / 2
}

func main() {
	fmt.Println(CalculateCircleArea(5))
	fmt.Println(CalculateRectangleArea(4, 4))
	fmt.Println(CalculateTriangleArea(5, 4))

}

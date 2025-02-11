package main

import "fmt"

func FindMaxAndMin(n ...int) (int, int) {
	if len(n) == 0 {
		return 0, 0
	}
	var max int = n[0]
	var min int = n[0]

	for _, n1 := range n {

		if max < n1 {
			max = n1
		}
		if min > n1 {
			min = n1
		}
	}
	return max, min
}
func main() {
	fmt.Println(FindMaxAndMin(1, 2, 3, 4, 5, 13, 5, 54, -2))
}

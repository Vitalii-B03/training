package main

import "fmt"

func main() {
	a := 5
	b := 10
	c := Dereference(&a)
	d := Sum(&b, &c)
	fmt.Println(c) // Output: 5
	fmt.Println(d) // Output: 15
}
func Dereference(n *int) int {
	d := *n
	return d
}
func Sum(n *int, c *int) int {
	sum := *n + *c
	return sum
}

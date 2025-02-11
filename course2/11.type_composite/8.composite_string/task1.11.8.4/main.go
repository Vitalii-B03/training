package main

import "fmt"

func concatStrings(xs ...string) string {
	var s string
	for _, str := range xs {
		s += str

	}
	return s
}
func main() {
	fmt.Println(concatStrings("hello", " ", "world"))
}

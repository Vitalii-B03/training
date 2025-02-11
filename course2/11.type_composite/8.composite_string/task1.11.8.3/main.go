package main

import "fmt"

func getBytes(s string) []byte {
	return []byte(s)
}
func getRunes(s string) []rune {
	return []rune(s)
}
func main() {
	s := "Привет, мир!"
	fmt.Println(getRunes(s))
	fmt.Println(getBytes(s))
}

package main

import "fmt"

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(SecondString())
	fmt.Println(ThirdString())
}

func HelloWorld() string {
	a := "Hello world!"
	return a
}
func SecondString() string {
	a := "This is second line!"
	return a
}
func ThirdString() string {
	a := "This is third line!"
	return a
}

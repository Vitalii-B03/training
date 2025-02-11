package main

func appendInt(xs *[]int, y ...int) {
	*xs = append(*xs, y...)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func UserInfo(name string, age int, cities ...string) string {
	StrAge := strconv.Itoa(age)
	return "Имя: " + name + ", возраст: " + StrAge + ", города: " + strings.Join(cities, ", ")
}
func main() {
	result := UserInfo("John", 21)
	fmt.Println(result)
}

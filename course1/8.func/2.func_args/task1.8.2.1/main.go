package main

import (
	"fmt"
	"strconv"
)

func UserInfo(name, city, phone string, age, weight int) string {
	StrAge := strconv.Itoa(age)
	StrWeight := strconv.Itoa(weight)
	res := "Имя: " + name + ", Город: " + city + ", Телефон: " + phone + ", Возраст: " + StrAge + ", Вес: " + StrWeight
	return res
}
func main() {
	fmt.Println(UserInfo("Jane", "Angeles", "987-654-3210", 25, 150))
}

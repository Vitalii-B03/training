package main

import (
	"fmt"
)

func getFilePermissions(flag int) string {
	var list = map[int]string{
		0: "-,-,-",
		1: "-,-,Execute",
		2: "-,Write,-",
		3: "-,Write,Execute",
		4: "Read,-,-",
		5: "Read,-,Execute",
		6: "Read,Write,-",
		7: "Read,Write,Execute",
	}
	if flag == 0 {
		c := "Owner:-,-,- Group:-,-,- Other:-,-,-"
		return c
	} else {
		owner := flag / 100
		group := (flag % 100) / 10
		other := flag % 10
		c := "Owner:" + list[owner] + " Group:" + list[group] + " Other:" + list[other]
		return c
	}
}
func main() {
	var a int
	fmt.Println("Введите значение прав доступа: ")
	fmt.Scan(&a)
	fmt.Println(getFilePermissions(a))

}

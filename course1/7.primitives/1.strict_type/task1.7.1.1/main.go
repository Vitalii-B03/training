package main

import "fmt"

func main() {
	var name, age, city string
	fmt.Println("Введите ваше имя, возраст и город рождения: ")
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Scanln(&city)
	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Город:", city)

}

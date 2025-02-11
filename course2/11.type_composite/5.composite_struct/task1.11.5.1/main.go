package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
)

type User struct {
	Name string
	Age  int
}

func getUsers() []User {
	var users []User
	for i := 0; i < 10; i++ {
		user := User{
			Name: gofakeit.Name(),
			Age:  gofakeit.Number(18, 60),
		}
		users = append(users, user)
	}
	return users
}
func preparePrint(users []User) string {
	var result string
	for _, user := range users {
		result += fmt.Sprintf("Имя: %s, Возраст: %d\n", user.Name, user.Age)
	}
	return result
}
func main() {
	users := getUsers()
	result := preparePrint(users)
	fmt.Println(result)
}

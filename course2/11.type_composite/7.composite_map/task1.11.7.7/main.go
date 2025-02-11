package main

import "fmt"

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	// Твой код для получения уникальных пользователей по никнейму
	var users1 []string
	for _, user := range users {
		users1 = append(users1, user.Nickname)
	}
	var userRes []User
	uniq := map[string]User{}
	for _, u := range users {
		if _, ok := uniq[u.Nickname]; !ok {
			uniq[u.Nickname] = u
		}

	}

	for _, u := range users1 {
		if _, ok := uniq[u]; ok {
			userRes = append(userRes, User{Nickname: u, Age: uniq[u].Age, Email: uniq[u].Email})
			delete(uniq, u)

		}

	}

	return userRes
}
func main() {
	us := User{Nickname: "John", Age: 25, Email: "john@gmail.com"}
	us1 := User{Nickname: "Jane", Age: 30, Email: "jane@example.com"}
	us2 := User{Nickname: "Alice", Age: 20, Email: "alice@example.com"}
	us3 := User{Nickname: "John", Age: 40, Email: "john@gmail.com"}
	users := getUniqueUsers([]User{us, us1, us2, us3})
	fmt.Println(users)
}

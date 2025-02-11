package main

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func NewUser(id int, opts ...UserOption) *User {
	user := &User{
		ID: id,
	}
	for _, opt := range opts {
		opt(user)
	}
	return user
}
func WithUsername(name string) UserOption {
	return func(user *User) {
		user.Username = name
	}
}
func WithEmail(mail string) UserOption {
	return func(user *User) {
		user.Email = mail
	}
}
func WithRole(role string) UserOption {
	return func(user *User) {
		user.Role = role
	}
}
func main() {
	user := NewUser(1, WithUsername("testuser"), WithEmail("testuser@example.com"), WithRole("admin"))
	fmt.Printf("User: %+v\n", user)
}

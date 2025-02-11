package main

import (
	"fmt"
	"github.com/icrowley/fake"
)

func main() {

	fmt.Println(GenerateFakeData())
}
func GenerateFakeData() string {
	name := fake.FirstName()
	lastname := fake.LastName()
	addres := fake.StreetAddress()
	phone := fake.Phone()
	email := fake.EmailAddress()
	a := "Name: " + name + " " + lastname + " " + "Address: " + addres + " " + "Phone: " + phone + " " + "Email: " + email
	return a
}

package main

import (
	"encoding/json"
	"fmt"
)

func getJSON(data []User) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Text string `json:"text"`
}

func main() {

	data := []User{
		{Name: "Jon", Age: 20, Comments: []Comment{{Text: "skjdvngjksnv"}}},
	}
	fmt.Println(getJSON(data))
}

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return nil, err
	}
	return db, nil
}
func CreateUserTable() error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    age INTEGER
)`)
	if err != nil {
		fmt.Println("Ошибка при создании таблицы:", err)
		return err
	}
	return nil
}

func InsertUser(user User) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO user ( name, age) VALUES ( ?, ?)", user.Name, user.Age)
	if err != nil {
		fmt.Println("Ошибка при вставке данных: ", err)
		return err
	}
	return nil
}

func SelectUser(id int) (User, error) {
	db, err := openDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close()
	row := db.QueryRow("SELECT id, name, age FROM user WHERE id = ?", id)
	var user User
	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		fmt.Println("Ошибка при выборке пользователя: ", err)
		return User{}, err
	}
	return user, nil
}

func UpdateUser(user User) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("UPDATE user SET name = ?, age = ?  WHERE id = ?", user.Name, user.Age, user.ID)
	if err != nil {
		fmt.Println("Ошибка при обновлении данных: ", err)
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		fmt.Println("Ошибка при удалении данных: ", err)
		return err
	}
	return nil
}
func main() {
	// Пример использования функций

	// Создание таблицы
	err := CreateUserTable()
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	// Вставка нового пользователя
	newUser := User{Name: "John Doe", Age: 30}
	err = InsertUser(newUser)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}

	// Выборка пользователя по ID
	user, err := SelectUser(1)
	if err != nil {
		log.Fatalf("Error selecting user: %v", err)
	}
	fmt.Printf("User: %v\n", user)

	// Обновление информации о пользователе
	user.Name = "Jane Doe"
	user.Age = 25
	err = UpdateUser(user)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}

	// Удаление пользователя
	err = DeleteUser(1)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    email TEXT
)`)
	if err != nil {
		fmt.Println("Ошибка при создании таблицы:", err)
		return err
	}
	return nil
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()
	query, args, err := PrepareQuery("insert", "user", user)
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		fmt.Println("Ошибка при вставке данных: ", err)
		return err
	}
	return nil
}

func SelectUser(id int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return User{}, err
	}
	defer db.Close()
	query, args, err := PrepareQuery("select", "user", User{ID: id})
	if err != nil {
		return User{}, err
	}
	var user User
	err = db.QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		fmt.Println("Ошибка при выборке пользователя: ", err)
		return User{}, err
	}
	return user, nil
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()
	query, args, err := PrepareQuery("update", "user", user)
	_, err = db.Exec(query, args...)
	if err != nil {
		fmt.Println("Ошибка при обновлении данных: ", err)
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()
	query, args, err := PrepareQuery("delete", "user", User{ID: id})
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		fmt.Println("Ошибка при удалении данных: ", err)
		return err
	}
	return nil
}
func PrepareQuery(operation string, table string, user User) (string, []interface{}, error) {
	switch operation {
	case "insert":
		query := squirrel.Insert(table).
			Columns("username", "email").
			Values(user.Username, user.Email)
		sqlQuery, args, err := query.ToSql()
		return sqlQuery, args, err
	case "select":
		query := squirrel.Select("id", "username", "email").
			From(table).
			Where(squirrel.Eq{"id": user.ID})
		sqlQureey, args, err := query.ToSql()
		return sqlQureey, args, err
	case "update":
		query := squirrel.Update(table).
			Set("email", user.Email).
			Set("username", user.Username).
			Where(squirrel.Eq{"id": user.ID})
		sqlQuery, args, err := query.ToSql()
		return sqlQuery, args, err
	case "delete":
		query := squirrel.Delete(table).
			Where(squirrel.Eq{"id": user.ID})
		sqlQuery, args, err := query.ToSql()
		return sqlQuery, args, err
	default:
		return "", nil, nil
	}
}
func main() {
	err := CreateUserTable()
	if err != nil {
		log.Fatal(err)
	}
	newUser := User{Username: "John Doe", Email: "rer@mail.ru"}
	err = InsertUser(newUser)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}
	user, err := SelectUser(1)
	if err != nil {
		log.Fatalf("Error selecting user: %v", err)
	}
	fmt.Printf("User: %v\n", user)

	user.Username = "Jane Doe"
	user.Email = "rer@mailgujutg.ru"
	err = UpdateUser(user)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	fmt.Printf("User: %v\n", user)
	err = DeleteUser(1)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
}

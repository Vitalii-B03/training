package main

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}
type Comment struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	UserID int    `json:"user_id"`
}

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "users.db?_timeout=20000") // увеличьте тайм-аут, если нужно
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return nil, err
	}
	return db, nil
}
func CreateUserTable() error {
	var err error
	db, err := getDBConnection()
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
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS comment (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    text TEXT,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES user(id)
)`)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func InsertUser(user User) error {
	db, err := getDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = prepareQuery("insert", "user", user).(squirrel.InsertBuilder).RunWith(db).Exec()
	if err != nil {
		return err
	}
	fmt.Println("Users добавлен")
	return nil
}
func SelectUser(userID int) (User, error) {
	db, err := getDBConnection()
	if err != nil {
		return User{}, err
	}
	defer db.Close()
	if err != nil {
		return User{}, err
	}
	var user User
	rows := prepareQuery("select", "user", User{ID: userID}).(squirrel.SelectBuilder).RunWith(db).QueryRow()
	err = rows.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
func UpdateUser(user User) error {
	db, err := getDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	// Подготовка запроса на обновление
	_, err = prepareQuery("update", "user", user).(squirrel.UpdateBuilder).RunWith(db).Exec()
	if err != nil {

		return err
	}
	fmt.Println("User обнавлен")
	return nil
}
func DeleteUser(userID int) error {
	db, err := getDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = prepareQuery("delete", "user", User{ID: userID}).(squirrel.DeleteBuilder).RunWith(db).Exec()
	if err != nil {
		return err
	}
	fmt.Println("User удален")
	return nil
}
func prepareQuery(operation string, table string, user User) interface{} {
	var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)
	switch operation {
	case "insert":
		return psql.Insert(table).
			Columns("name", "age").
			Values(user.Name, user.Age)
	case "select":
		return psql.Select("id, name, age").
			From(table).
			Where(squirrel.Eq{"id": user.ID})
	case "update":
		return psql.Update(table).
			Set("name", user.Name).
			Set("age", user.Age).
			Where(squirrel.Eq{"id": user.ID})
	case "delete":
		return psql.Delete("user").
			Where(squirrel.Eq{"id": user.ID})
	default:
		return nil
	}

}
func main() {
	// Создание таблицы
	err := CreateUserTable()
	if err != nil {
		log.Fatal(err)
	}

	//Добавление нового пользователя
	Newuser := User{Name: "Joj", Age: 25}
	err = InsertUser(Newuser)
	if err != nil {
		log.Fatal(err)
	}
	//Выбор пользователя
	user, err := SelectUser(1)
	if err != nil {
		log.Fatalf("Error selecting user: %v", err)
	}
	fmt.Printf("User: %v\n", user)
	// Обновление пользователя
	user.Name = "Jane Doe"
	user.Age = 30
	err = UpdateUser(user)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	fmt.Printf("User: %v\n", user)
	// Удаление пользователя
	err = DeleteUser(1)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
}

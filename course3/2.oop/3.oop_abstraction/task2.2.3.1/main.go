package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"reflect"
	"strings"
)

// Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

type Tabler interface {
	TableName() string
}

func (u User) TableName() string {
	return "user"
}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}
type SQLiteGenerator struct{}

func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	var colums []string

	e := reflect.ValueOf(table).Elem()
	for i := 0; i < e.NumField(); i++ {
		dbField := e.Type().Field(i).Tag.Get("db_field")
		dbType := e.Type().Field(i).Tag.Get("db_type")
		colums = append(colums, fmt.Sprintf("%s %s", dbField, dbType))
	}
	return fmt.Sprintf("CREATE TABLE %s (%s);", table.TableName(), strings.Join(colums, ", "))
}

func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	var colums []string
	var values []string
	e := reflect.ValueOf(model).Elem()
	for i := 0; i < e.NumField(); i++ {
		dbField := e.Type().Field(i).Tag.Get("db_field")
		values = append(values, e.Field(i).String())
		colums = append(colums, fmt.Sprintf("%s", dbField))
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", model.TableName(), strings.Join(colums, ", "), strings.Join(values, ", "))
}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenerateFakeUser() User
}
type GoFakeitGenerator struct{}

func (g *GoFakeitGenerator) GenerateFakeUser() User {
	return User{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}

}
func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}

	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)

	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}
}

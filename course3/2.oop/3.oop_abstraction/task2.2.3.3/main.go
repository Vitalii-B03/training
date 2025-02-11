package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"reflect"
	"strings"
)

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}
type Tabler interface {
	TableName() string
}
type SQLGenerator interface {
	CreateTableSQL(t Tabler) string
	CreateInsertSQL(t Tabler) string
}
type FakeDataGenerator interface {
	GenerateFakeUser() User
}

func (s *SQLiteGenerator) CreateTableSQL(t Tabler) string {
	var colums []string

	e := reflect.ValueOf(t).Elem()
	for i := 0; i < e.NumField(); i++ {
		dbField := e.Type().Field(i).Tag.Get("db_field")
		dbType := e.Type().Field(i).Tag.Get("db_type")
		colums = append(colums, fmt.Sprintf("%s %s", dbField, dbType))
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", t.TableName(), strings.Join(colums, ", "))
}
func (s *SQLiteGenerator) CreateInsertSQL(t Tabler) string {
	var colums []string
	var values []string
	e := reflect.ValueOf(t).Elem()
	for i := 0; i < e.NumField(); i++ {
		dbField := e.Type().Field(i).Tag.Get("db_field")
		colums = append(colums, fmt.Sprintf("%s", dbField))
		field := e.Field(i)
		var value string
		switch field.Kind() {
		case reflect.Int:
			value = fmt.Sprintf("%d", field.Int())
		case reflect.String:
			value = fmt.Sprintf("'%s'", field.String())
		default:
			value = "NULL"
		}
		values = append(values, value)
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", t.TableName(), strings.Join(colums, ", "), strings.Join(values, ", "))
}

func (u User) TableName() string {
	return "users"
}

type SQLiteGenerator struct {
}
type GoFakeitGenerator struct {
}

func (g *GoFakeitGenerator) GenerateFakeUser() User {
	u := User{
		ID:        gofakeit.Number(1000, 9999),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
	return u
}
func GenerateUserInserts(value int) []string {
	fakeDataGenerator := &GoFakeitGenerator{}
	sqlGenerator := &SQLiteGenerator{}
	var result []string
	for i := 0; i < value; i++ {

		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		result = append(result, query)

	}
	return result
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

	queries := GenerateUserInserts(34)
	for _, query := range queries {
		fmt.Println(query)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int    `db_field:"id" db_type:"INTEGER PRIMARY KEY AUTOINCREMENT"`
	FirstName string `db_field:"first_name" db_type:"TEXT"`
	LastName  string `db_field:"last_name" db_type:"TEXT"`
	Email     string `db_field:"email" db_type:"TEXT UNIQUE"`
}

type Migrator struct {
	db           *sql.DB
	sqlGenerator SQLGenerator
}

func (u *User) TableName() string {
	return "users"
}

type Tabler interface {
	TableName() string
}

func (m *Migrator) Migrate(models ...Tabler) error {
	for _, model := range models {
		query := m.sqlGenerator.CreateTableSQL(model)
		_, err := m.db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to create table for model %v: %v", model.TableName(), err)
		}
	}
	return nil
}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}
type SQLiteGenerator struct {
}

func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	e := reflect.ValueOf(table).Elem()
	var colums []string
	for i := 0; i < e.NumField(); i++ {
		dbField := e.Type().Field(i).Tag.Get("db_field")
		dbType := e.Type().Field(i).Tag.Get("db_type")
		colums = append(colums, fmt.Sprintf("%s %s", dbField, dbType))
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", table.TableName(), strings.Join(colums, ", "))
}
func (s *SQLiteGenerator) CreateInsertSQL(table Tabler) string {
	return ""
}
func NewMigrator(db *sql.DB, sqlGenerator SQLGenerator) *Migrator {
	return &Migrator{
		db:           db,
		sqlGenerator: sqlGenerator,
	}
}

// Основная функция
func main() {
	// Подключение к SQLite БД
	db, err := sql.Open("sqlite3", "file:my_database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Создание мигратора с использованием вашего SQLGenerator
	sqlGenerator := &SQLiteGenerator{}
	migrator := NewMigrator(db, sqlGenerator)

	// Миграция таблицы User
	if err := migrator.Migrate(&User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
	fmt.Println("Migrator успешно сщздан")
}

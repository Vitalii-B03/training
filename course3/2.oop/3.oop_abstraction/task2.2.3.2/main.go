package main

// Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}
type SQLiteGenerator struct {
}
type GoFakeitGenerator struct {
}

type Tabler interface {
	TableName() string
}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(i interface{}) string
	CreateInsertSQL(i interface{}) string
}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenerateFakeUser()
}

func (s *SQLiteGenerator) CreateTableSQL(i interface{}) string {
	//TODO implement me
	panic("implement me")
}

func (s *SQLiteGenerator) CreateInsertSQL(i interface{}) string {
	//TODO implement me
	panic("implement me")
}

func (g *GoFakeitGenerator) GenerateFakeUser() string {
	//TODO implement me
	panic("implement me")
}

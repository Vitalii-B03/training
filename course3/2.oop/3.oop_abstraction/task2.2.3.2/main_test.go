package main

import "testing"

func TestCreateTableSQL(t *testing.T) {
	sqlGenerator := SQLiteGenerator{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Паника не вызвана")
		} else if r != "implement me" {
			t.Errorf("Текст ошибки нк совпадает, ожидаемый `implement me`, а полученный %v", r)
		}
	}()
	sqlGenerator.CreateTableSQL(&User{})
}
func TestCreateInsertSQL(t *testing.T) {
	sqlGenerator := SQLiteGenerator{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Паника не вызвана")
		} else if r != "implement me" {
			t.Errorf("Текст ошибки нк совпадает, ожидаемый `implement me`, а полученный %v", r)
		}
	}()
	sqlGenerator.CreateInsertSQL(&User{})
}
func TestGoFakeitGenerator(t *testing.T) {
	generator := GoFakeitGenerator{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Паника не вызвана")
		} else if r != "implement me" {
			t.Errorf("Текст ошибки нк совпадает, ожидаемый `implement me`, а полученный %v", r)
		}

	}()
	generator.GenerateFakeUser()

}

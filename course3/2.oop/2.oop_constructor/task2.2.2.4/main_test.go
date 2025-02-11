package main

import (
	"bytes"
	"os"
	"testing"
)

func TestFileLogger_Log(t *testing.T) {
	file, err := os.CreateTemp("", "test_log_*.txt")
	if err != nil {
		t.Fatalf("ошибка создания временного файла %v", err)
	}
	defer os.Remove(file.Name())
	fileLogger := &FileLogger{file: file}
	message := "test"
	fileLogger.Log(message)

	file.Seek(0, 0)
	var buf bytes.Buffer
	_, err = buf.ReadFrom(file)
	if err != nil {
		t.Fatalf("Ошибка чтения временного файла %v", err)
	}
	actual := buf.String()
	expected := message + "\n"
	if actual != expected {
		t.Errorf("Проверяемое значение %q, а полученное %q", expected, actual)
	}

}
func TestNewLogSystem_WithLogger(t *testing.T) {
	file, err := os.CreateTemp("", "test_log_*.txt")
	if err != nil {
		t.Fatalf("ошибка создания временного файла %v", err)
	}
	defer os.Remove(file.Name())

	fileLogger := &FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))
	if logSystem.logger == nil {
		t.Errorf("Логгер не установлен")
	}
	message := "test"
	logSystem.Log(message)

	file.Seek(0, 0)
	var buf bytes.Buffer
	_, err = buf.ReadFrom(file)
	if err != nil {
		t.Fatalf("Ошибка чтения временного файла %v", err)
	}
	actual := buf.String()
	expected := message + "\n"
	if actual != expected {
		t.Errorf("Проверяемое значение %q, а полученное %q", expected, actual)
	}
}

type TestLogger struct {
	Messages []string
}

func (t *TestLogger) Log(message string) {
	t.Messages = append(t.Messages, message)
}
func TestLogSystem_Log(t *testing.T) {
	testLogger := &TestLogger{}
	logSystem := NewLogSystem(WithLogger(testLogger))

	message := "test"
	logSystem.Log(message)
	if len(testLogger.Messages) != 1 {
		t.Errorf("Ожидаем 1 сообщение, а получили %d", len(testLogger.Messages))
	}
	if testLogger.Messages[0] != message {
		t.Errorf("Вместо сообщенрия %q, получили %q", message, testLogger.Messages[0])
	}
}

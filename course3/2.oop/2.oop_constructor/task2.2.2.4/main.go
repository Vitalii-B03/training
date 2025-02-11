package main

import (
	"fmt"
	"io"
	"os"
)

// Logger interface
type Logger interface {
	Log(message string)
}

// FileLogger struct
type FileLogger struct {
	file *os.File
}

func (fl *FileLogger) Log(message string) {
	_, err := fl.file.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println("Ошибка записи в файл", err)
	}
}

type LogSystem struct {
	logger Logger
}

func (ls *LogSystem) Log(message string) {
	ls.logger.Log(message)
}

// ConsoleLogger struct
type ConsoleLogger struct {
	out io.ReadWriter
}

// LogOption functional option type
type LogOption func(*LogSystem)

func WithLogger(logger Logger) LogOption {
	return func(ls *LogSystem) {
		ls.logger = logger
	}
}
func NewLogSystem(opts ...LogOption) *LogSystem {
	logSystem := &LogSystem{}
	for _, option := range opts {
		option(logSystem)
	}
	return logSystem
}
func main() {
	file, _ := os.Create("log.txt")
	defer file.Close()


	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(&fileLogger))

	logSystem.Log("Hello, world!")
}

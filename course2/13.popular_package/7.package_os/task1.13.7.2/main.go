package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func WriteFile(reader io.Reader, fd io.Writer) error {
	file, ok := fd.(*os.File)
	if !ok {
		return errors.New("file is not a os.File")
	}
	filePath := file.Name()
	dirPath := filePath[:len(filePath)-len(file.Name())]
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}
	_, err := io.Copy(file, reader)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func main() {
	filePath := "course1/13.popular_package/7.package_os/task1.13.7.2/file.txt"

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer file.Close() // отложенная функция закрытия дескриптора файла

	err = WriteFile(strings.NewReader("Hello, World!"), file)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
	}
}

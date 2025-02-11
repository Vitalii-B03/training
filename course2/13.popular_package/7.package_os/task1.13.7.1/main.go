package main

import (
	"fmt"
	"os"
)

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	dir := filePath[:len(filePath)-len("file.txt")]
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	err := WriteFile("/path/to/file.txt", []byte("hello world"), os.FileMode(0644))
	if err != nil {
		fmt.Println("Error:", err)
	}
}

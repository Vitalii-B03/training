package main

import (
	"bytes"
	"fmt"
)

func getDataString(b *bytes.Buffer) string {
	buf := make([]byte, b.Len())
	r := bytes.NewReader(b.Bytes())
	r.Read(buf)
	return string(buf)
}
func main() {
	// Create a new buffer
	buffer := bytes.NewBufferString("Hello, World!")

	// Call the getDataString function
	result := getDataString(buffer)

	// Check if the result matches the expected output
	expected := "Hello, World!"
	if result != expected {
		panic(fmt.Sprintf("Expected %s, but got %s", expected, result))
	}
}

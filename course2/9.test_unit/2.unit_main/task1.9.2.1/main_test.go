package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	oldStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()

	os.Stdout = oldStdout

	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	expected := "Hello, world!\n"
	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}

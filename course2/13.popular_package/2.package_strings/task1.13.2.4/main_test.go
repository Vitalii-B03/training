package main

import (
	"strings"
	"testing"
)

func TestGenerateActivationKey(t *testing.T) {
	key := generateActivationKey()
	expectedLen := 19
	if len(key) != expectedLen {
		t.Errorf("Expected length of %d, got %d", expectedLen, len(key))
	}
	valChars := "1234567890QWERTYUIOPLKJHGFDSAZXCVBNM-"
	for _, char := range key {
		if !strings.Contains(valChars, string(char)) {
			t.Errorf("Generated string %s does not contain the string '%v'", valChars, string(char))
		}
	}
	for i, char := range key {
		if (i+1)%5 == 0 && char != '-' {
			t.Errorf("Generated string %s does not contain the string '%v'", valChars, string(char))
		}
	}
}

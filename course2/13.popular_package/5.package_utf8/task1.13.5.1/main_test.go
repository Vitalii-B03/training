package main

import "testing"

func TestCountUniqueUTF8Chars(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{"test1", "qqwweerrtt", 5},
		{"test2", "1q2w3e", 6},
		{"test3", "1q2wee", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUniqueUTF8Chars(tt.text); got != tt.want {
				t.Errorf("countUniqueUTF8Chars() = %v, want %v", got, tt.want)
			}
		})
	}

}
func BenchmarkCountUniqueUTF8Chars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countUniqueUTF8Chars("qqwweerrtt")
	}

}

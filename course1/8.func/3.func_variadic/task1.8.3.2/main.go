package main

import (
	"fmt"
	"strings"
)

func ConcatenateStrings(sep string, str ...string) string {
	var res1 []string
	var res2 []string

	for str1 := range str {
		if str1%2 == 0 {
			res1 = append(res1, str[str1])

		} else {
			res2 = append(res2, str[str1])
		}

	}
	ResStr := "even: " + strings.Join(res1, sep) + ", odd: " + strings.Join(res2, sep)
	ResStr = strings.TrimSuffix(ResStr, sep)
	
	return ResStr

}
func main() {
	fmt.Println(ConcatenateStrings("_"))
}

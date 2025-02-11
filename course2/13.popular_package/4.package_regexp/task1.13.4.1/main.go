package main

import (
	"fmt"
	"regexp"
)

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`@.*\.com$`)
	if re.MatchString(email) {
		return true
	} else {
		return false
	}
}
func main() {
	email := "test@example.co"
	valid := isValidEmail(email)
	if valid {
		fmt.Printf("%s является валидным email-адресом\n", email)
	} else {
		fmt.Printf("%s не является валидным email-адресом\n", email)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func generateActivationKey() string {
	letters := "QWERTYUIOPASDFGHJKLZXCVBNM0123456789"
	resbox := []string{}
	rand.Seed(time.Now().UnixNano())
	var xres string
	for i := 0; i < 4; i++ {
		xres = ""
		for j := 0; j < 4; j++ {

			cl := rand.Intn(2)
			if cl == 0 {
				xres += string(letters[rand.Intn(len(letters))])
			} else {
				xres += strconv.Itoa(rand.Intn(10))
			}

		}
		resbox = append(resbox, xres)

	}
	return strings.Join(resbox, "-")
}
func main() {
	fmt.Println(generateActivationKey())

}

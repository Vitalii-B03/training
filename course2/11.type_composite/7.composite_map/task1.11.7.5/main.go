package main

import (
	"fmt"
	"strings"
)

func filterSentence(sentence string, filter map[string]bool) string {
	var str []string
	var strRes string
	var mapStr = make(map[string]bool)
	str = strings.Split(sentence, " ")
	for _, word := range str {
		mapStr[word] = filter[word]
	}
	for _, word := range str {
		if mapStr[word] == false {
			strRes += word + " "
		}
	}
	strRes = strRes[:len(strRes)-1]
	return strRes
}
func main() {
	sentence := "ipsum this is elit a test"
	filter := map[string]bool{"ipsum": true, "elit": true}
	fmt.Println(filterSentence(sentence, filter))
}

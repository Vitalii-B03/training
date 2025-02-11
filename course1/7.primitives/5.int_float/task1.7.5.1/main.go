package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func binaryStringToFloat(binary string) float32 {
	var number uint32
	number1, _ := strconv.ParseUint(binary, 2, 32)
	number = uint32(number1)
	floatNumber := *(*float32)(unsafe.Pointer(&number))
	return floatNumber
}
func main() {
	fmt.Println(binaryStringToFloat(("00111110001000000000000000000000")))
}

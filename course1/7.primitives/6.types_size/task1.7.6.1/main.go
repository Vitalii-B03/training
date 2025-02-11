package main

import (
	"fmt"
	"unsafe"
)

func sizeOfBool(b bool) int {
	c := unsafe.Sizeof(b)
	return int(c)
}
func sizeOfInt(n int) int {
	c := unsafe.Sizeof(n)
	return int(c)
}
func sizeOfInt8(n int8) int {
	c := unsafe.Sizeof(n)
	return int(c)
}
func sizeOfInt16(n int16) int {
	c := unsafe.Sizeof(n)
	return int(c)
}
func sizeOfInt32(n int32) int {
	c := unsafe.Sizeof(n)
	return int(c)
}
func sizeOfInt64(n int64) int {
	c := unsafe.Sizeof(n)
	return int(c)
}
func sizeOfUint(n uint) int {
	c := unsafe.Sizeof(n)
	return int(c)
}
func sizeOfUint8(n uint8) int {
	c := unsafe.Sizeof(n)
	return int(c)
}
func main() {
	var b bool
	var n int
	var n8 int8
	var n16 int16
	var n32 int32
	var n64 int64
	var u uint
	var u8 uint8
	fmt.Printf("boll size: %d\n", sizeOfBool(b))
	fmt.Printf("int size: %d\n", sizeOfInt(n))
	fmt.Printf("int8 size: %d\n", sizeOfInt8(n8))
	fmt.Printf("int16 size: %d\n", sizeOfInt16(n16))
	fmt.Printf("int32 size: %d\n", sizeOfInt32(n32))
	fmt.Printf("int64 size: %d\n", sizeOfInt64(n64))
	fmt.Printf("uint size: %d\n", sizeOfUint(u))
	fmt.Printf("uint8 size: %d\n", sizeOfUint8(u8))
}

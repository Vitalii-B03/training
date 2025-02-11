package main

import (
	"fmt"
	"github.com/ptflp/gocolors"
)

func main() {
	a := "Hello, world!"
	fmt.Println(ColorizeCustom(a, 200, 100, 50))

}

func ColorizeRed(a string) string {
	return gocolors.Colorize(gocolors.Red, a)
}
func ColorizeGreen(a string) string {
	return gocolors.Colorize(gocolors.Green, a)
}
func ColorizeBlue(a string) string {
	return gocolors.Colorize(gocolors.Blue, a)
}
func ColorizeYellow(a string) string {
	return gocolors.Colorize(gocolors.Yellow, a)
}
func ColorizeWhite(a string) string {
	return gocolors.Colorize(gocolors.White, a)
}
func ColorizeMagenta(a string) string {
	return gocolors.Colorize(gocolors.Magenta, a)
}
func ColorizeCyan(a string) string {
	return gocolors.Colorize(gocolors.Cyan, a)
}
func ColorizeCustom(a string, r, g, b int) string {
	a1 := gocolors.RGB(r, g, b)
	return gocolors.Colorize(a1, a)
}

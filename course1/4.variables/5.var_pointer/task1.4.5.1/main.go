package main

func main() {
	var a int
	var b float64
	var c string
	var d bool
	changeInt(&a)
	changeFloat(&b)
	changeString(&c)
	changeBool(&d)

}
func changeInt(a *int) {
	*a = 20
}
func changeFloat(b *float64) {
	*b = 6.28
}
func changeString(c *string) {
	*c = "Goodbye, world!"
}
func changeBool(d *bool) {
	*d = false
}

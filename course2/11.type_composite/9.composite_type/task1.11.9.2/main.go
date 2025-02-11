package main

import "fmt"

type LgTV struct {
	status bool
	model  string
}

type SamsungTV struct {
	status bool
	model  string
}
type TVer interface {
	switchOFF() bool
	switchON() bool
	GetStatus() bool
	GetModel() string
}

func (tvs *SamsungTV) switchOFF() bool {
	tvs.status = false
	return true
}
func (tvs *SamsungTV) switchOn() bool {
	tvs.status = true
	return true
}
func (tvs *SamsungTV) GetStatus() bool {
	return tvs.status
}
func (tvs *SamsungTV) GetModel() string {
	return tvs.model
}
func (tvs *SamsungTV) SamsungHub() string {
	return "SamsungHub"
}
func (tvg *LgTV) LGHub() string {
	return "LGHub"
}
func (tvg *LgTV) switchOFF() bool {
	tvg.status = false
	return true
}
func (tvg *LgTV) switchOn() bool {
	tvg.status = true
	return true
}
func (tvg *LgTV) GetStatus() bool {
	return tvg.status
}
func (tvg *LgTV) GetModel() string {
	return tvg.model
}
func main() {
	tvs := &SamsungTV{
		status: true,
		model:  "Samsung XL-100500",
	}
	fmt.Println(tvs.GetStatus())  // true
	fmt.Println(tvs.GetModel())   // Samsung XL-100500
	fmt.Println(tvs.SamsungHub()) // SamsungHub
	fmt.Println(tvs.GetStatus())  // true
	fmt.Println(tvs.switchOn())   // true
	fmt.Println(tvs.GetStatus())  // true
	fmt.Println(tvs.switchOFF())  // true
	fmt.Println(tvs.GetStatus())  // false

	tvg := &LgTV{
		status: true,
		model:  "LG XL-100500",
	}
	fmt.Println(tvg.GetStatus()) // true
	fmt.Println(tvg.GetModel())  // Samsung XL-100500
	fmt.Println(tvg.LGHub())     // SamsungHub
	fmt.Println(tvg.GetStatus()) // true
	fmt.Println(tvg.switchOn())  // true
	fmt.Println(tvg.GetStatus()) // true
	fmt.Println(tvg.switchOFF()) // true
	fmt.Println(tvg.GetStatus()) // false
}

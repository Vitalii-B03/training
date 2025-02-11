package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		timeNow := time.Now()
		dateOnly := timeNow.Format("2006-01-02") //02-01-2006
		timeN := timeNow.Format("15:04:05")
		time.Sleep(1 * time.Second)
		fmt.Print("\033[H\033[2J") // очистка терминала
		fmt.Println("Текущее время:", timeN)
		fmt.Println("Текущая дата:", dateOnly)
	}
}

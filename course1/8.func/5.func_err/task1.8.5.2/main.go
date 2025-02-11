package main

import "fmt"

func CheckDiscount(price, discount float64) (float64, error) {
	if discount > 50 {
		str := "Скидка не может быть больше 50%"
		return 0, fmt.Errorf("%s\n", str)
	} else {
		return price - (price * discount / 100), nil
	}
}
func main() {
	result, err := CheckDiscount(500, 50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

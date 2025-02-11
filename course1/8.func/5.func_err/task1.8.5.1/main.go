package main

import (
	"fmt"
	"strconv"
)

func CalculatePercentageChange(initialValue, finalValue string) (float64, error) {
	val, err := strconv.ParseFloat(initialValue, 64)
	if err != nil {
		return 0, fmt.Errorf("Недопустимое начальное значение: %v", err)

	}
	final, err2 := strconv.ParseFloat(finalValue, 64)
	if err2 != nil {
		return 0, fmt.Errorf("Недопустимое конечное значение: %v", err)
	}
	if val == 0 {
		return 0, nil
	}

	return ((final - val) / val) * 100, nil
}
func main() {
	result, err := CalculatePercentageChange("0", "2")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("%.2f%%\n", result)
	}
}

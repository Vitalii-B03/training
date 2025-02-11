package main

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	a := "2.1"
	b := "2"

	fmt.Println(DecimalEqual(a, b))
}
func DecimalSum(a, b string) (string, error) {
	a1, err1 := decimal.NewFromString(a)
	b1, err2 := decimal.NewFromString(b)
	if err1 != nil || err2 != nil {
		return "", errors.New("Не корректные данные")

	}
	sum := a1.Add(b1)
	return sum.String(), nil
}
func DecimalSubtract(a, b string) (string, error) {
	a1, err1 := decimal.NewFromString(a)
	b1, err2 := decimal.NewFromString(b)
	if err1 != nil || err2 != nil {
		return "", errors.New("Не корректные данные")

	}
	sub := a1.Sub(b1)
	return sub.String(), nil
}
func DecimalMultiply(a, b string) (string, error) {
	a1, err1 := decimal.NewFromString(a)
	b1, err2 := decimal.NewFromString(b)
	if err1 != nil || err2 != nil {
		return "", errors.New("Не корректные данные")

	}
	mul := a1.Mul(b1)
	return mul.String(), nil
}
func DecimalDivide(a, b string) (string, error) {
	a1, err1 := decimal.NewFromString(a)
	b1, err2 := decimal.NewFromString(b)
	if err1 != nil || err2 != nil {
		return "", errors.New("Не корректные данные")
	}
	if b == "0" {
		return "", errors.New("На 0 делить нельзя")

	}
	div := a1.Div(b1)
	return div.String(), nil
}
func DecimalRound(a string, precision int32) (string, error) {
	a1, err := decimal.NewFromString(a)
	if err != nil {
		return "", errors.New("Не корректные данные")
	}
	return a1.Round(precision).String(), nil

}
func DecimalGreaterThan(a, b string) (bool, error) {
	a1, err1 := decimal.NewFromString(a)
	b1, err2 := decimal.NewFromString(b)
	if err1 != nil || err2 != nil {
		return true, errors.New("Не корректные данные")
	}
	div := decimal.Max(a1, b1)
	if div == a1 {
		return true, nil
	} else {
		return false, nil
	}
}
func DecimalLessThan(a, b string) (bool, error) {
	a1, err1 := decimal.NewFromString(a)
	b1, err2 := decimal.NewFromString(b)
	if err1 != nil || err2 != nil {
		return true, errors.New("Не корректные данные")
	}
	div := decimal.Max(a1, b1)
	if div == a1 {
		return false, nil
	} else {
		return true, nil
	}
}
func DecimalEqual(a, b string) (bool, error) {
	a1, err1 := decimal.NewFromString(a)
	b1, err2 := decimal.NewFromString(b)
	if err1 != nil || err2 != nil {
		return true, errors.New("Не корректные данные")
	}
	div := a1.Cmp(b1)

	if div == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

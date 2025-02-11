package main

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) error
}
type CreditCard struct {
	balance float64
}

func (c *CreditCard) Pay(amount float64) error {
	if amount > c.balance {
		return errors.New("недостаточный баланс")
	} else if amount < 0 {
		return errors.New("недопустимая сумма платежа")
	}
	c.balance = c.balance - amount
	fmt.Printf("Оплачено %.2f с помощью кредитной карты\n", amount)
	return nil
}

type Bitcoin struct {
	balance float64
}

func (b *Bitcoin) Pay(amount float64) error {
	if amount > b.balance {
		return errors.New("недостаточный баланс")
	} else if amount < 0 {
		return errors.New("недопустимая сумма платежа")
	}
	b.balance = b.balance - amount
	fmt.Printf("Оплачено %.2f с помощью биткоина\n", amount)
	return nil
}
func ProcessPayment(p PaymentMethod, amount float64) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Println("Не удалось обработать платеж:", err)
	}
}

func main() {
	cc := &CreditCard{balance: 500.00}
	btc := &Bitcoin{balance: 2.00}

	ProcessPayment(cc, 200.00)
	ProcessPayment(btc, 1.00)
}

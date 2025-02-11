package main

import (
	"errors"
	"fmt"
)

type Accounter interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}
type CurrentAccount struct {
	balance float64
}

func (c *CurrentAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("negative amount")
	}
	c.balance += amount
	return nil

}
func (c *CurrentAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return errors.New("Отрицательная сумма")
	}
	if c.balance < amount {
		return errors.New("Недостаточно средств")
	}
	c.balance -= amount
	return nil
}
func (c *CurrentAccount) Balance() float64 {
	return c.balance
}

type SavingsAccount struct {
	balance float64
}

func (s *SavingsAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Отрицательное сумма")
	}
	s.balance += amount
	return nil

}
func (s *SavingsAccount) Withdraw(amount float64) error {
	if s.balance-amount < 500 {
		return errors.New("Баланс карты менее 500!!!")
	} else if amount < 0 {
		return errors.New("Недостаточно средств или отрицательная сумма")
	} else {

		s.balance -= amount
		return nil
	}
}
func (s *SavingsAccount) Balance() float64 {
	return s.balance
}

func ProcessAccount(a Accounter) {
	a.Deposit(500)
	a.Withdraw(200)
	fmt.Printf("Balance: %.2f\n", a.Balance())
}

func main() {
	c := &CurrentAccount{}
	s := &SavingsAccount{}
	ProcessAccount(c)
	ProcessAccount(s)
}

package main

import (
	"errors"
	"fmt"
)

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}
type DineInOrder struct {
	orderDetails map[string]int
}

func (d *DineInOrder) AddItem(item string, quantity int) error {
	if quantity <= 0 {
		return errors.New("Некоректно количество в заказе.")
	}
	d.orderDetails[item] = quantity
	fmt.Printf("Заказ изменен: %v\n", d.GetOrderDetails())
	return nil
}
func (d *DineInOrder) RemoveItem(item string) error {
	if _, ok := d.orderDetails[item]; ok {
		delete(d.orderDetails, item)
		fmt.Printf("Заказ изменен: %v\n", d.GetOrderDetails())
		return nil
	}
	return errors.New("Заказ не найден")

}
func (d *DineInOrder) GetOrderDetails() map[string]int {
	return d.orderDetails
}

type TakeAwayOrder struct {
	orderDetails map[string]int
}

func (t *TakeAwayOrder) AddItem(item string, quantity int) error {
	if quantity <= 0 {
		return errors.New("Некоректно количество в заказе.")
	}
	t.orderDetails[item] = quantity
	fmt.Printf("Заказ изменен: %v\n", t.GetOrderDetails())
	return nil
}
func (t *TakeAwayOrder) RemoveItem(item string) error {
	if _, ok := t.orderDetails[item]; ok {
		delete(t.orderDetails, item)
		fmt.Printf("Заказ изменен: %v\n", t.GetOrderDetails())
		return nil
	}
	return errors.New("Нет такого закуаза")
}
func (t *TakeAwayOrder) GetOrderDetails() map[string]int {
	return t.orderDetails
}
func ManageOrder(o Order) {
	o.AddItem("Pizza", 2)
	o.AddItem("Burger", 1)
	o.RemoveItem("Pizza")
	fmt.Println(o.GetOrderDetails())
}

func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}

	ManageOrder(dineIn)
	ManageOrder(takeAway)
}

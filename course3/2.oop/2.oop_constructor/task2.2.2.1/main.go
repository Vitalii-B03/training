package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type OrderOption func(*Order)

func NewOrder(i int, options ...OrderOption) *Order {
	order := &Order{
		ID: i,
	}
	for _, option := range options {
		option(order)
	}
	return order
}
func WithOrderDate(date time.Time) OrderOption {
	return func(o *Order) {
		o.OrderDate = date
	}
}
func WithCustomerID(customerID string) OrderOption {
	return func(o *Order) {
		o.CustomerID = customerID
	}
}
func WithItems(items []string) OrderOption {
	return func(o *Order) {
		o.Items = items

	}
}

func main() {
	order := NewOrder(2,
		WithCustomerID("123"),
		WithItems([]string{"item1", "item2"}),
		WithOrderDate(time.Now()))

	fmt.Printf("Order: %+v\n", order)
}

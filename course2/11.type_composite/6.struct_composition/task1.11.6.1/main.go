package main

import "fmt"

type Dish struct {
	Name  string
	Price float64
}
type Order struct {
	Dishes []Dish
	Total  float64
}

func (order *Order) AddDish(dish Dish) {
	order.Dishes = append(order.Dishes, dish)
}
func (order *Order) CalculateTotal() {
	order.Total = 0
	for _, dish := range order.Dishes {
		order.Total += dish.Price
	}
}
func (order *Order) RemoveDish(dish Dish) {
	for i, dish2 := range order.Dishes {
		if dish == dish2 {
			order.Dishes = append(order.Dishes[:i], order.Dishes[i+1:]...)
		}

	}
}
func main() {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}
	order.AddDish(dish1)
	order.AddDish(dish2)
	order.CalculateTotal()
	fmt.Println(order.Dishes, order.Total)
	order.RemoveDish(dish1)
	order.CalculateTotal()
	fmt.Println(order.Dishes, order.Total)
}

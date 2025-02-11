package main

import "fmt"

type Animal struct {
	Type string
	Name string
	Age  int
}

func getAnimals() []Animal {
	var animals []Animal
	animal1 := Animal{
		Type: "Кот",
		Name: "Барсик",
		Age:  9,
	}
	animal2 := Animal{
		Type: "Собака",
		Name: "Бобик",
		Age:  6,
	}
	animal3 := Animal{
		Type: "Попугай",
		Name: "Кеша",
		Age:  3,
	}
	animals = append(animals, animal1, animal2, animal3)
	return animals
}
func preparePrint(animals []Animal) string {
	var animalStr string
	for _, animal := range animals {
		animalStr += fmt.Sprintf("Тип: %s, Имя: %s, Возраст: %d\n", animal.Type, animal.Name, animal.Age)
	}
	return animalStr
}
func main() {
	animals := getAnimals()
	animalStr := preparePrint(animals)
	fmt.Print(animalStr)
}

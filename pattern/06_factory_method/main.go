package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

	Применимость:

	- Когда необходимо делегировать ответственность за создание объектов подклассам
	- Когда система должна быть независимой от способа создания, композиции и представления объектов
	- Когда создание объекта является деталями реализации и должно быть скрыто от клиентского кода

	Плюсы:

	- Нет тесной связи между создателем и конкретными продуктами.
	- Принцип единой ответственности. Вы можете переместить код создания продукта в одно место программы, что упростит поддержку кода.
	- Принцип открытости/закрытости. Вы можете вводить в программу новые виды продуктов, не нарушая существующий клиентский код.

	Минусы:

	- Усложнение кода

	Примеры использования:

	- Библиотеки для работы с базами данных
	- Фреймворки для создания графических интерфейсов
*/

// Products

type Product interface {
	GetName() string
}

type ProductA struct{}

var _ Product = (*ProductA)(nil)

func (p *ProductA) GetName() string {
	return "Product A"
}

type ProductB struct{}

var _ Product = (*ProductB)(nil)

func (p *ProductB) GetName() string {
	return "Product B"
}

// Creators

type Creator interface {
	CreateProduct() Product
}

type ConcreteCreatorA struct{}

var _ Creator = (*ConcreteCreatorA)(nil)

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ProductA{}
}

type ConcreteCreatorB struct{}

var _ Creator = (*ConcreteCreatorB)(nil)

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ProductB{}
}

func main() {
	creatorA := ConcreteCreatorA{}
	productA := creatorA.CreateProduct()
	fmt.Printf("%s\n", productA.GetName())

	creatorB := ConcreteCreatorB{}
	productB := creatorB.CreateProduct()
	fmt.Printf("%s\n", productB.GetName())
}

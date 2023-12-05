package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

	Factory method (фабричный метод) — это шаблон порождающего проектирования, который предоставляет интерфейс для создания объектов в суперклассе, но позволяет подклассам изменять тип создаваемых объектов.

	Применимость:

	1. Когда необходимо делегировать ответственность за создание объектов подклассам: Если существует несколько классов, реализующих общий интерфейс, но способы создания этих объектов различны, то Factory Method позволяет каждому подклассу решать, какой именно объект создавать.
	2. Когда система должна быть независимой от способа создания, композиции и представления объектов: Паттерн позволяет избежать привязки клиентского кода к конкретным классам объектов, которые создаются, что делает систему более гибкой и расширяемой.
	3. Когда создание объекта является деталями реализации и должно быть скрыто от клиентского кода: Factory Method помогает сокрыть детали создания объектов, предоставляя абстрактный метод для их создания в интерфейсе создателя.

	Плюсы:

	1. Гибкость и расширяемость: Паттерн позволяет добавлять новые продукты или изменять существующие, не изменяя клиентский код.
	2. Снижение зависимости: Клиентский код зависит от абстрактного создателя и интерфейса продукта, что позволяет избежать привязки к конкретным классам.
	3. Повышение уровня абстракции: Factory Method помогает разделить создание объекта и его использование, повышая уровень абстракции в коде.
*/

// Products

type ProductInterface interface {
	GetName() string
}

type ProductA struct{}

var _ ProductInterface = (*ProductA)(nil)

func (p *ProductA) GetName() string {
	return "Product A"
}

type ProductB struct{}

var _ ProductInterface = (*ProductB)(nil)

func (p *ProductB) GetName() string {
	return "Product B"
}

// Creators

type CreatorInterface interface {
	CreateProduct() ProductInterface
}

type CreatorA struct{}

var _ CreatorInterface = (*CreatorA)(nil)

func (c *CreatorA) CreateProduct() ProductInterface {
	return &ProductA{}
}

type CreatorB struct{}

var _ CreatorInterface = (*CreatorB)(nil)

func (c *CreatorB) CreateProduct() ProductInterface {
	return &ProductB{}
}

func main() {
	creatorA := CreatorA{}
	productA := creatorA.CreateProduct()
	fmt.Printf("%s\n", productA.GetName())

	creatorB := CreatorB{}
	productB := creatorB.CreateProduct()
	fmt.Printf("%s\n", productB.GetName())
}

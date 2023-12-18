package main

import (
	"log"
)

/*
	Реализовать паттерн «строитель».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

	Builder — это порождающий шаблон проектирования, который позволяет шаг за шагом создавать сложные объекты.

	Применимость:

	1. Сложные объекты: Используйте Builder, когда создание сложного объекта состоит из множества шагов, и вам нужно разделить процесс конструирования на отдельные этапы.
	2. Различные конфигурации: Если у вас есть несколько вариантов конфигурации объекта, и вы хотите избежать "телескопического конструктора" (множества конструкторов с разным числом параметров).
	3. Изоляция конструирования: Используйте Builder, чтобы изолировать сложные шаги конструирования от самого объекта, делая код более управляемым и читаемым.

	Плюсы:

	1. Разделение ответственностей: Каждый строитель отвечает за конкретный шаг конструирования, что делает код более поддерживаемым и гибким.
	2. Избегание телескопического конструктора: Помогает избежать создания множества конструкторов с разным числом параметров, что может сделать код трудным для понимания и использования.
	3. Поддержка различных представлений продукта: Вы можете использовать один и тот же строитель для создания различных представлений одного продукта.

	Минусы:

	1. Усложнение кода: Использование Builder может привести к увеличению количества классов и усложнению кода, особенно если объект, который вы строите, не очень сложен.
	2. Дублирование кода: Если у вас много различных строителей, может возникнуть дублирование кода, связанного с конструированием.
*/

type Product struct {
	Part1 string
	Part2 string
}

type BuilderInterface interface {
	BuildPart1()
	BuildPart2()

	GetProduct() Product
}

type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) BuildPart1() {
	b.product.Part1 = "Part 1 built"
}

func (b *ConcreteBuilder) BuildPart2() {
	b.product.Part2 = "Part 2 built"
}

func (b *ConcreteBuilder) GetProduct() Product {
	return b.product
}

type Director struct {
	builder BuilderInterface
}

func NewDirector(builder BuilderInterface) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() Product {
	d.builder.BuildPart1()
	d.builder.BuildPart2()

	return d.builder.GetProduct()
}

func main() {
	builder := &ConcreteBuilder{}
	director := NewDirector(builder)

	product := director.Construct()

	log.Printf("%s\n", product.Part1)
	log.Printf("%s\n", product.Part2)
}

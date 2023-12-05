package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

	Visitor — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции, не изменяя классы объектов, над которыми эти операции могут выполняться.

	Применимость:

	1. Когда у вас есть множество объектов с разными типами, и вы хотите выполнить различные операции в зависимости от типа объекта.
	2. Когда добавление новых операций к объектам становится проблемой из-за их разнообразия.
	3. Когда структура объектов часто изменяется, но операции над ними остаются постоянными.

	Плюсы:

	1. Разделение обязанностей: Visitor позволяет разделить операции над объектами от их структуры, что упрощает добавление новых операций без изменения самих объектов.
	2. Легкость добавления новых типов: Добавление нового типа объекта (новой реализации интерфейса) и новой операции (нового метода в посетителе) является относительно простой процедурой.
	3. Централизованный код: Операции посетителя централизованы в одном месте, что делает код более поддерживаемым и понятным.

	Минусы:

	1. Сложность добавления новых классов: Добавление нового класса (нового типа объекта) требует изменения всех Visitor ( посетителей ), что может быть неудобным.
	2. Нарушение инкапсуляции: Visitor может иметь доступ к приватным членам объекта, что может привести к нарушению инкапсуляции.
	3. Усложнение структуры программы: Внедрение Visitor ( поситителя ) может привести к более сложной структуре программы из-за большего числа интерфейсов и классов.
*/

type VisitableInterface interface {
	Accept(VisitorInterface)
}

type ElementA struct {
	Name string
}

var _ VisitableInterface = (*ElementA)(nil)

func (c *ElementA) Accept(v VisitorInterface) {
	v.VisitElementA(c)
}

type ElementB struct {
	Number int
}

var _ VisitableInterface = (*ElementB)(nil)

func (c *ElementB) Accept(v VisitorInterface) {
	v.VisitElementB(c)
}

type VisitorInterface interface {
	VisitElementA(*ElementA)
	VisitElementB(*ElementB)
}

type Visitor struct{}

var _ VisitorInterface = (*Visitor)(nil)

func (v *Visitor) VisitElementA(elementA *ElementA) {
	fmt.Printf("%s\n", elementA.Name)
}

func (v *Visitor) VisitElementB(elementB *ElementB) {
	fmt.Printf("%d\n", elementB.Number)
}

func main() {
	elementA := &ElementA{Name: "name"}
	elementB := &ElementB{Number: 1}

	visitor := &Visitor{}

	elementA.Accept(visitor)
	elementB.Accept(visitor)
}

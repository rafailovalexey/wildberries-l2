package main

import (
	"log"
)

/*
	Реализовать паттерн «цепочка вызовов».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

	Chain of responsibility — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

	Применимость:

	1. Когда у вас есть несколько объектов, способных обработать запрос: Используется, когда запрос может быть обработан одним из нескольких объектов, и вы хотите, чтобы обработчики были независимыми друг от друга.
	2. Когда вы не знаете заранее, какой объект должен обработать запрос: Позволяет динамически устанавливать порядок обработчиков и добавлять новые обработчики без изменения кода клиента.
	3. Когда обработчики должны обрабатывать запросы последовательно: Когда нужно убедиться, что запрос обрабатывается только одним обработчиком, и порядок обработки важен.

	Плюсы:

	1. Разделение обязанностей: Каждый обработчик отвечает за свою часть логики, что способствует разделению обязанностей.
	2. Гибкость и расширяемость: Легко добавлять новые обработчики или изменять порядок существующих без изменения клиентского кода.
	3. Уменьшение связанности: Клиентский код не привязан к конкретным классам обработчиков, что уменьшает связанность между клиентом и обработчиками.

	Минусы:

	1. Гарантии обработки: Нет гарантий, что запрос будет обработан. Возможно, он дойдет до конца цепочки без обработки.
	2. Сложность отладки: Сложно отслеживать и отлаживать, поскольку запрос может пройти через несколько обработчиков до достижения конечного.
*/

type HandlerInterface interface {
	HandleRequest(request int) bool
	SetNext(handler HandlerInterface)
}

type BaseHandler struct {
	NextHandler HandlerInterface
}

func (h *BaseHandler) SetNext(handler HandlerInterface) {
	h.NextHandler = handler
}

type HandlerA struct {
	BaseHandler
}

func (h *HandlerA) HandleRequest(request int) bool {
	if request <= 10 {
		log.Printf("handler a %d\n", request)

		return true
	}

	if h.NextHandler != nil {
		return h.NextHandler.HandleRequest(request)
	}

	return false
}

type HandlerB struct {
	BaseHandler
}

func (h *HandlerB) HandleRequest(request int) bool {
	if request > 10 && request <= 20 {
		log.Printf("handler b %d\n", request)

		return true
	}

	if h.NextHandler != nil {
		return h.NextHandler.HandleRequest(request)
	}

	return false
}

func main() {
	handlerA := &HandlerA{}
	handlerB := &HandlerB{}

	handlerA.SetNext(handlerB)

	requests := []int{5, 10, 15, 20, 25}

	for _, request := range requests {
		if !handlerA.HandleRequest(request) {
			log.Printf("no handler can process the request %d\n", request)
		}
	}
}

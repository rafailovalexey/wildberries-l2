package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

	Strategy — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

	Применимость:

	1. Необходимость выбора алгоритма на лету: Когда у вас есть несколько вариантов алгоритмов, которые могут быть использованы динамически во время выполнения программы.
	2. Исключение дублирования кода: Когда у вас есть ряд алгоритмов, которые могут содержать одинаковый код, и использование стратегии позволяет избежать дублирования.
	3. Инкапсуляция алгоритмов: Когда вы хотите инкапсулировать алгоритмы и сделать их независимыми от контекста, чтобы легко добавлять, удалять или изменять алгоритмы без влияния на остальную часть кода.

	Плюсы:

	1. Избегание условных операторов: Strategy позволяет избежать длинных последовательностей условных операторов, заменяя их отдельными стратегиями.
	2. Улучшение читаемости кода: Код становится более читаемым и понятным, так как каждая Strategy находится в своем классе.
	3. Легкость расширения: Новые Strategy могут быть легко добавлены без изменения существующего кода.

	Минусы:

	1. Увеличение числа классов: Каждая Strategy обычно требует свой собственный класс, что может привести к увеличению числа классов в программе.
	2. Сложность выбора стратегии: В некоторых случаях может возникнуть сложность выбора подходящей стратегии.
	3. Накладные расходы на передачу данных: Если контексту часто приходится передавать данные в Strategy, это может привести к дополнительным накладным расходам.
*/

type PaymentStrategyInterface interface {
	Pay(amount float64) string
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("paid %.2f using credit card", amount)
}

type QiwiPayment struct{}

func (p *QiwiPayment) Pay(amount float64) string {
	return fmt.Sprintf("paid %.2f using qiwi", amount)
}

type ShoppingCart struct {
	paymentStrategy PaymentStrategyInterface
}

func (s *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategyInterface) {
	s.paymentStrategy = strategy
}

func (s *ShoppingCart) MakePayment(amount float64) string {
	if s.paymentStrategy == nil {
		return "no payment strategy set"
	}

	return s.paymentStrategy.Pay(amount)
}

func main() {
	cart := &ShoppingCart{}

	cart.SetPaymentStrategy(&CreditCardPayment{})
	fmt.Printf("%s\n", cart.MakePayment(100.50))

	cart.SetPaymentStrategy(&QiwiPayment{})
	fmt.Printf("%s\n", cart.MakePayment(75.25))
}

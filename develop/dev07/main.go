package main

import (
	"fmt"
	"time"
)

/*
	=== Or channel ===

	Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
	Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
	однако иногда неизвестно общее число done каналов, с которыми вы работаете в runtime.
	В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

	Определение функции:
	var or func(channels ...<- chan interface{}) <- chan interface{}

	Пример использования функции:

	sig := func(after time.Duration) <- chan interface{} {
		c := make(chan interface{})

		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()

	<-or (
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf(“done after %v\n”, time.Since(start))
*/

type ApplicationInterface interface {
	or(channels ...<-chan interface{}) <-chan interface{}
}

type Application struct{}

var _ ApplicationInterface = (*Application)(nil)

func main() {
	application := &Application{}

	signal := func(after time.Duration) <-chan interface{} {
		channel := make(chan interface{})

		go func() {
			defer close(channel)

			time.Sleep(after)
		}()

		return channel
	}

	start := time.Now()

	<-application.or(
		signal(2*time.Hour),
		signal(5*time.Minute),
		signal(5*time.Millisecond),
		signal(1*time.Hour),
		signal(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}

func (a *Application) or(channels ...<-chan interface{}) <-chan interface{} {
	output := make(chan interface{})

	for _, channel := range channels {
		go func(channel <-chan interface{}) {
			select {
			case <-channel:
				close(output)
			}
		}(channel)
	}

	return output
}

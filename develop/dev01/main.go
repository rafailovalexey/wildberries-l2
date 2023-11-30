package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

/*
	=== Базовая задача ===

	Создать программу, печатающую точное время с использованием NTP библиотеки. Инициализировать как go module.
	Использовать библиотеку https://github.com/beevik/ntp.
	Написать программу, печатающую текущее время / точное время с использованием этой библиотеки.

	Программа должна быть оформлена с использованием как go module.
	Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
	Программа должна проходить проверки go vet и golint.
*/

type ApplicationInterface interface {
	GetTimeNtp() (time.Time, error)
}

type Application struct{}

var _ ApplicationInterface = (*Application)(nil)

func main() {
	application := &Application{}

	ntpTime, err := application.GetTimeNtp()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка при получении времени: %v\n", err)

		os.Exit(1)
	}

	fmt.Printf("точное время (с использованием NTP): %s\n", ntpTime.Format(time.RFC3339))
}

func (a *Application) GetTimeNtp() (time.Time, error) {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		return time.Time{}, err
	}

	return ntpTime, nil
}

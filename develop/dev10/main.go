package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

/*
	=== Утилита telnet ===

	Реализовать примитивный telnet клиент:
	Примеры вызовов:
	go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

	Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
	После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
	Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

	При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
	При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type FlagsInterface interface {
	InitializeFlags()
}

type Flags struct {
	Timeout time.Duration
}

var _ FlagsInterface = (*Flags)(nil)

func (f *Flags) InitializeFlags() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout for connection")

	flag.Parse()

	f.Timeout = *timeout
}

type ApplicationInterface interface {
	Connection(address string, flags *Flags)
}

type Application struct{}

var _ ApplicationInterface = (*Application)(nil)

func main() {
	flags := &Flags{}
	application := &Application{}

	flags.InitializeFlags()

	host := flag.Arg(0)
	port := flag.Arg(1)

	if host == "" || port == "" {
		fmt.Printf("usage: [--timeout=<timeout>] <host> <port>\n")

		os.Exit(1)
	}

	address := fmt.Sprintf("%s:%s", host, port)

	application.Connection(address, flags)
}

func (a *Application) Connection(address string, flags *Flags) {
	connection, err := net.DialTimeout("tcp", address, flags.Timeout)

	if err != nil {
		fmt.Printf("error connecting to the server %v\n", err)

		os.Exit(1)
	}

	defer connection.Close()

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt)

	go func() {
		buffer := make([]byte, 1024)

		for {
			length, err := connection.Read(buffer)

			if err != nil {
				fmt.Printf("error reading from the server %v\n", err)

				break
			}

			fmt.Print(string(buffer[:length]))
		}
	}()

	buffer := make([]byte, 1024)

	for {
		length, err := os.Stdin.Read(buffer)

		if err != nil {
			fmt.Printf("error reading from STDIN %v\n", err)

			break
		}

		_, err = connection.Write(buffer[:length])

		if err != nil {
			fmt.Printf("error writing to the server %v\n", err)

			break
		}
	}

	select {
	case <-done:
		fmt.Printf("connection closed by user\n")
	case <-time.After(flags.Timeout):
		fmt.Printf("timeout reached, connection closed\n")
	}
}

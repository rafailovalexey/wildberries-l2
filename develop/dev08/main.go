package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
	=== Взаимодействие с ОС ===

	Необходимо реализовать собственный шелл

	встроенные команды: cd/pwd/echo/kill/ps
	поддержать fork/exec команды
	конвейер на пайпах

	Реализовать утилиту netcat (nc) клиент
	принимать данные из stdin и отправлять в соединение (tcp/udp)
	Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Command struct {
	name   string
	handle func([]string) (bool, error)
}

type Commands = []Command

type ApplicationInterface interface {
	GetCommands() map[string]Command
}

type Application struct{}

var _ ApplicationInterface = (*Application)(nil)

func main() {
	application := &Application{}
	commands := application.GetCommands()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("$ ")

		scanner.Scan()

		input := scanner.Text()

		arguments := strings.Fields(input)

		if command, isExist := commands[arguments[0]]; isExist {
			exit, err := command.handle(arguments)

			if err != nil {
				fmt.Printf("%v\n", err)
			}

			if exit {
				break
			}
		}

		if _, isExist := commands[arguments[0]]; !isExist {
			exit, err := commands["default"].handle(arguments)

			if err != nil {
				fmt.Printf("%v\n", err)
			}

			if exit {
				break
			}
		}
	}
}

func (a *Application) GetCommands() map[string]Command {
	commands := &Commands{
		Command{
			name: "cd",
			handle: func(arguments []string) (bool, error) {
				if len(arguments) < 2 {
					return false, fmt.Errorf("использование: cd <directory>")
				}

				err := os.Chdir(arguments[1])

				if err != nil {
					return false, fmt.Errorf("ошибка при изменении директории %v", err)
				}

				return false, nil
			},
		},
		Command{
			name: "pwd",
			handle: func(arguments []string) (bool, error) {
				directory, err := os.Getwd()

				if err != nil {
					return false, fmt.Errorf("ошибка при получении текущей директории %v", err)
				}

				fmt.Printf("%s\n", directory)

				return false, nil
			},
		},
		Command{
			name: "echo",
			handle: func(arguments []string) (bool, error) {
				fmt.Printf("%s\n", strings.Join(arguments[1:], " "))

				return false, nil
			},
		},
		Command{
			name: "kill",
			handle: func(arguments []string) (bool, error) {
				if len(arguments) < 2 {
					return false, fmt.Errorf("использование: kill <pid>")
				}

				cmd := exec.Command("kill", arguments[1:]...)

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Run(); err != nil {
					return false, fmt.Errorf("ошибка выполнения команды %v", err)
				}

				return false, nil
			},
		},
		Command{
			name: "ps",
			handle: func(arguments []string) (bool, error) {
				cmd := exec.Command("ps")

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Run(); err != nil {
					return false, fmt.Errorf("ошибка выполнения команды %v", err)
				}

				return false, nil
			},
		},
		Command{
			name: "exit",
			handle: func(arguments []string) (bool, error) {
				fmt.Printf("выход из терминала\n")

				return true, nil
			},
		},
		Command{
			name: "default",
			handle: func(arguments []string) (bool, error) {
				cmd := exec.Command(arguments[0], arguments[1:]...)

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Run(); err != nil {
					return false, fmt.Errorf("ошибка выполнения команды %v", err)
				}

				return false, nil
			},
		},
	}

	dictionary := make(map[string]Command, len(*commands))

	for _, command := range *commands {
		dictionary[command.name] = command
	}

	return dictionary
}

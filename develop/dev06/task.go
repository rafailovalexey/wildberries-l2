package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
	=== Утилита cut ===

	Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

	Поддержать флаги:
	-f - "fields" - выбрать поля (колонки)
	-d - "delimiter" - использовать другой разделитель
	-s - "separated" - только строки с разделителем

	Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type FlagsInterface interface {
	InitializeFlags()
}

type Flags struct {
	Fields    string
	Delimiter string
	Separated bool
}

var _ FlagsInterface = (*Flags)(nil)

func (f *Flags) InitializeFlags() {
	fields := flag.String("f", "", "")
	delimiter := flag.String("d", "	", "")
	separated := flag.Bool("s", false, "")

	flag.Parse()

	f.Fields = *fields
	f.Delimiter = *delimiter
	f.Separated = *separated
}

type ApplicationInterface interface {
	ParseNumberColumns(string) int
}

type Application struct{}

var _ ApplicationInterface = (*Application)(nil)

func main() {
	flags := &Flags{}
	application := &Application{}

	flags.InitializeFlags()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if flags.Separated && !strings.Contains(line, flags.Delimiter) {
			fmt.Printf("\n")

			continue
		}

		fields := strings.Split(line, flags.Delimiter)

		if flags.Fields != "" {
			temporary := make([]string, 0)

			field := strings.Split(flags.Fields, ",")

			for _, v := range field {
				index := application.ParseNumberColumns(v)

				if index > 0 && index <= len(fields) {
					temporary = append(temporary, fields[index-1])
				}
			}

			fmt.Printf("%s\n", strings.Join(temporary, flags.Delimiter))

			fmt.Printf("\n")

			continue
		}

		fmt.Printf("%s\n", strings.Join(fields, flags.Delimiter))

		fmt.Printf("\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("ошибка чтения стандартного ввода %v\n", err)

		os.Exit(1)
	}
}

func (a *Application) ParseNumberColumns(str string) int {
	var result int

	_, err := fmt.Sscanf(str, "%d", &result)

	if err != nil {
		return 0
	}

	return result
}

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
	delimiter := flag.String("d", "", "")
	separated := flag.Bool("s", false, "")

	flag.Parse()

	f.Fields = *fields
	f.Delimiter = *delimiter
	f.Separated = *separated
}

type ApplicationInterface interface{}

type Application struct{}

var _ ApplicationInterface = (*Application)(nil)

func main() {
	flags := &Flags{}

	flags.InitializeFlags()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if flags.Separated && !strings.Contains(line, flags.Delimiter) {
			continue
		}

		fieldsList := strings.Split(line, flags.Delimiter)

		var selectedFields []string

		if flags.Fields != "" {
			fieldIndices := strings.Split(flags.Fields, ",")

			for _, indexStr := range fieldIndices {
				index := parseInt(indexStr)

				if index > 0 && index <= len(fieldsList) {
					selectedFields = append(selectedFields, fieldsList[index-1])
				}
			}
		} else {
			selectedFields = fieldsList
		}

		fmt.Println(strings.Join(selectedFields, flags.Delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка чтения стандартного ввода:", err)
	}
}

func parseInt(s string) int {
	var result int

	_, err := fmt.Sscanf(s, "%d", &result)

	if err != nil {
		return 0
	}

	return result
}

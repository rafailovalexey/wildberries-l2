package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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
	CutInputStrings(*bufio.Scanner, *Flags)
	ParseNumberColumns(string) int
}

type Application struct{}

var _ ApplicationInterface = (*Application)(nil)

func main() {
	flags := &Flags{}
	application := &Application{}

	flags.InitializeFlags()

	scanner := bufio.NewScanner(os.Stdin)

	application.CutInputStrings(scanner, flags)

	if err := scanner.Err(); err != nil {
		log.Printf("error reading standard input %v\n", err)

		os.Exit(1)
	}
}

func (a *Application) CutInputStrings(scanner *bufio.Scanner, flags *Flags) {
	for scanner.Scan() {
		line := scanner.Text()

		if flags.Separated && !strings.Contains(line, flags.Delimiter) {
			continue
		}

		fields := strings.Split(line, flags.Delimiter)

		if flags.Fields != "" {
			temporary := make([]string, 0)

			for _, v := range fields {
				index := a.ParseNumberColumns(v)

				if index > 0 && index <= len(fields) {
					temporary = append(temporary, fields[index-1])
				}
			}

			if len(temporary) != 0 {
				log.Printf("%s\n", strings.Join(temporary, flags.Delimiter))
			}
		}
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

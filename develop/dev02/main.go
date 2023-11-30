package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
	=== Задача на распаковку ===

	Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
		- "a4bc2d5e" => "aaaabccddddde"
		- "abcd" => "abcd"
		- "45" => "" (некорректная строка)
		- "" => ""
	Дополнительное задание: поддержка escape - последовательностей
		- qwe\4\5 => qwe45 (*)
		- qwe\45 => qwe44444 (*)
		- qwe\\5 => qwe\\\\\ (*)

	В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

	Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type ApplicationInterface interface {
	UnpackString(string) string
}

type Application struct{}

var _ ApplicationInterface = (*Application)(nil)

func main() {
	application := &Application{}

	fmt.Println(application.UnpackString("a4bc2d5e") == "aaaabccddddde")
	fmt.Println(application.UnpackString("abcd") == "abcd")
	fmt.Println(application.UnpackString("45") == "")
	fmt.Println(application.UnpackString("") == "")

	fmt.Println(application.UnpackString("qwe\\4\\5") == "qwe45")
	fmt.Println(application.UnpackString("qwe\\45") == "qwe44444")
	fmt.Println(application.UnpackString("qwe\\\\5") == "qwe\\\\\\\\\\")
}

func (a *Application) UnpackString(str string) string {
	current := ""
	factor := 1
	escape := false

	result := &strings.Builder{}

	for _, r := range str {
		if r == 92 {
			if escape {
				current = string(r)
				factor = 1
				escape = false

				result.WriteString(strings.Repeat(current, factor))

				continue
			}

			escape = true

			continue
		}

		if unicode.IsDigit(r) {
			if escape {
				current = string(r)
				factor = 1
				escape = false

				result.WriteString(strings.Repeat(current, factor))

				continue
			}

			number, err := strconv.Atoi(string(r))

			if err != nil {
				factor = 1

				continue
			}

			factor = number - 1
		}

		if unicode.IsLetter(r) {
			current = string(r)
			factor = 1
			escape = false
		}

		result.WriteString(strings.Repeat(current, factor))
	}

	return result.String()
}

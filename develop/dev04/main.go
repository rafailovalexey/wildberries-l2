package main

import (
	"log"
	"sort"
	"strings"
)

/*
	=== Поиск анаграмм по словарю ===

	Напишите функцию поиска всех множеств анаграмм по словарю.
	Например:
	'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
	'листок', 'слиток' и 'столик' - другому.

	Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
	Выходные данные: Ссылка на мапу множеств анаграмм.
	Ключ - первое встретившееся в словаре слово из множества
	Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
	Множества из одного элемента не должны попасть в результат.
	Все слова должны быть приведены к нижнему регистру.
	В результате каждое слово должно встречаться только один раз.

	Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type ApplicationInterface interface {
	FindAnagrams(words []string) map[string][]string
}

type Application struct{}

var _ ApplicationInterface = (*Application)(nil)

func main() {
	application := &Application{}

	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagrams := application.FindAnagrams(words)

	for key, value := range anagrams {
		log.Printf("many anagrams for %s %s\n", key, strings.Join(value, ", "))
	}
}

func (a *Application) FindAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, value := range words {
		word := strings.ToLower(value)
		chars := strings.Split(word, "")

		sort.Strings(chars)

		sorted := strings.Join(chars, "")

		anagrams[sorted] = append(anagrams[sorted], word)
	}

	for key, value := range anagrams {
		if len(value) <= 1 {
			delete(anagrams, key)
		} else {
			sort.Strings(value)

			anagrams[key] = value
		}
	}

	return anagrams
}

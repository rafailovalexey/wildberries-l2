package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
	"unicode"
)

/*
	=== Утилита sort ===

	Отсортировать строки (man sort)
	Основное

	Поддержать ключи

	-k — указание колонки для сортировки
	-n — сортировать по числовому значению
	-r — сортировать в обратном порядке +
	-u — не выводить повторяющиеся строки +

	Дополнительное

	Поддержать ключи

	-M — сортировать по названию месяца
	-b — игнорировать хвостовые пробелы +
	-c — проверять отсортированы ли данные
	-h — сортировать по числовому значению с учётом суффиксов

	Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	inputFile := flag.String("i", "", "")
	outputFile := flag.String("o", "", "")

	//key := flag.Bool("k", false, "")
	//numeric := flag.String("n", false, "")
	//reverse := flag.Bool("r", false, "")
	//unique := flag.Bool("u", false, "")
	//
	//month := flag.Bool("M", false, "")
	//ignore := flag.Bool("b", false, "")
	//check := flag.Bool("c", false, "")
	//suffix := flag.Bool("h", false, "")

	flag.Parse()

	if *inputFile == "" {
		fmt.Printf("%v\n", "укажите файл для чтения")

		os.Exit(1)
	}

	inputFilepath, err := getFilepath(*inputFile)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	data, err := getFileData(inputFilepath)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	if *outputFile == "" {
		fmt.Printf("%v\n", "укажите файл для записи")

		os.Exit(1)
	}

	outputFilepath, err := getFilepath(*outputFile)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	sortedStrings := getStringsWithArguments(data, true, true, true)

	err = writeFileData(outputFilepath, sortedStrings)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}
}

func getWorkDirectory() (string, error) {
	pwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return pwd, nil
}

func getFilepath(file string) (string, error) {
	wd, err := getWorkDirectory()

	if err != nil {
		return "", err
	}

	filepath := path.Join(wd, file)

	return filepath, nil
}

func getFileData(filepath string) ([]string, error) {
	data := make([]string, 0, 10)

	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		data = append(data, line)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func writeFileData(filepath string, data []string) error {
	file, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, v := range data {
		newline := getNewline()

		_, err = writer.WriteString(v + newline)

		if err != nil {
			return err
		}
	}

	err = writer.Flush()

	if err != nil {
		return err
	}

	return nil
}

func getNewline() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}

	return "\n"
}

func getStringsWithArguments(data []string, unique bool, reverse bool, trailing bool) []string {
	temporary := make([]string, len(data))

	copy(temporary, data)

	if unique {
		temporary = make([]string, 0, len(data))

		dictionary := make(map[string]struct{}, len(data))

		for _, v := range data {
			dictionary[v] = struct{}{}
		}

		for v := range dictionary {
			temporary = append(temporary, v)
		}
	}

	if trailing {
		for i, v := range temporary {
			temporary[i] = strings.TrimRightFunc(v, unicode.IsSpace)
		}
	}

	if reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(temporary)))
	}

	if !reverse {
		sort.Strings(temporary)
	}

	return temporary
}

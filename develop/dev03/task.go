package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"sort"
	"strconv"
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
	-r — сортировать в обратном порядке
	-u — не выводить повторяющиеся строки

	Дополнительное

	Поддержать ключи

	-M — сортировать по названию месяца
	-b — игнорировать хвостовые пробелы
	-c — проверять отсортированы ли данные
	-h — сортировать по числовому значению с учётом суффиксов

	Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	inputFile := flag.String("i", "", "")
	outputFile := flag.String("o", "", "")

	column := flag.Int("k", 0, "")
	numeric := flag.Bool("n", false, "")
	reverse := flag.Bool("r", false, "")
	unique := flag.Bool("u", false, "")

	month := flag.Bool("M", false, "")
	ignore := flag.Bool("b", false, "")
	check := flag.Bool("c", false, "")
	suffix := flag.Bool("h", false, "")

	flag.Parse()

	if *inputFile == "" {
		fmt.Printf("%v\n", "укажите файл для чтения")

		os.Exit(1)
	}

	inputFilepath, err := GetFilepath(*inputFile)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	data, err := GetFileData(inputFilepath)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	if *outputFile == "" {
		fmt.Printf("%v\n", "укажите файл для записи")

		os.Exit(1)
	}

	outputFilepath, err := GetFilepath(*outputFile)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	result, err := GetSortedStringsWithArguments(
		data,
		*column,
		*numeric,
		*reverse,
		*unique,
		*month,
		*ignore,
		*check,
		*suffix,
	)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	err = WriteFileData(outputFilepath, result)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}
}

func GetFilepath(file string) (string, error) {
	wd, err := getWorkDirectory()

	if err != nil {
		return "", err
	}

	filepath := path.Join(wd, file)

	return filepath, nil
}

func GetFileData(filepath string) ([]string, error) {
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

func WriteFileData(filepath string, data []string) error {
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

func GetSortedStringsWithArguments(
	data []string,
	column int,
	numeric bool,
	reverse bool,
	unique bool,
	month bool,
	ignore bool,
	check bool,
	suffix bool,
) ([]string, error) {
	temporary := make([]string, len(data))

	copy(temporary, data)

	if ignore {
		temporary = GetStringsWithRemoveTrailingSpace(temporary)
	}

	function := func(i, j int) bool {
		if month {
			return getMonthValue(temporary[i]) < getMonthValue(temporary[j])
		}

		if column > 0 {
			return getSortColumnKey(temporary[i], column) < getSortColumnKey(temporary[j], column)
		}

		return temporary[i] < temporary[j]
	}

	if numeric {
		function = func(i, j int) bool {
			return getNumericValue(temporary[i]) < getNumericValue(temporary[j])
		}
	}

	if suffix {
		function = func(i, j int) bool {
			valueI, suffixI := getNumericAndSuffix(temporary[i])
			valueJ, suffixJ := getNumericAndSuffix(temporary[j])

			if valueI < valueJ {
				return true
			}

			if valueI > valueJ {
				return false
			}

			return suffixI < suffixJ
		}
	}

	if unique {
		temporary = GetUniqueStrings(temporary)
	}

	sort.SliceStable(temporary, function)

	if reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(temporary)))
	}

	if check && !CheckSortedStrings(temporary) {
		return nil, errors.New("данные не отсортированы")
	}

	return temporary, nil
}

func GetStringsWithRemoveTrailingSpace(data []string) []string {
	temporary := make([]string, len(data))

	copy(temporary, data)

	for i, v := range temporary {
		temporary[i] = strings.TrimRightFunc(v, unicode.IsSpace)
	}

	return temporary
}

func GetUniqueStrings(data []string) []string {
	temporary := make([]string, 0, len(data))
	dictionary := make(map[string]struct{}, len(data))

	for _, v := range data {
		if _, isExist := dictionary[v]; !isExist {
			dictionary[v] = struct{}{}
			temporary = append(temporary, v)
		}
	}

	return temporary
}

func CheckSortedStrings(data []string) bool {
	check := sort.IsSorted(sort.StringSlice(data))

	return check
}

func getWorkDirectory() (string, error) {
	pwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return pwd, nil
}

func getNewline() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}

	return "\n"
}

func getNumericValue(s string) int {
	num, err := strconv.Atoi(s)

	if err != nil {
		return 999999999
	}

	return num
}

func getNumericAndSuffix(input string) (int, string) {
	match := regexp.MustCompile(`^(\d+)([a-zA-Z]*)$`).FindStringSubmatch(input)

	if len(match) != 3 {
		return 9999999999, input
	}

	value, err := strconv.Atoi(match[1])

	if err != nil {
		return 9999999999, input
	}

	suffix := match[2]

	return value, suffix
}

func getSortColumnKey(s string, column int) string {
	columns := strings.Fields(s)

	if column > 0 && column <= len(columns) {
		return columns[column-1]
	}

	return s
}

func getMonthValue(month string) int {
	months := map[string]int{
		"january":   1,
		"february":  2,
		"march":     3,
		"april":     4,
		"may":       5,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
	}

	return months[strings.ToLower(month)]
}

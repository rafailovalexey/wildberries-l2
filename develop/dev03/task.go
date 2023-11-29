package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"reflect"
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

	-k — указание колонки для сортировки +
	-n — сортировать по числовому значению +
	-r — сортировать в обратном порядке +
	-u — не выводить повторяющиеся строки +

	Дополнительное

	Поддержать ключи

	-M — сортировать по названию месяца +
	-b — игнорировать хвостовые пробелы +
	-c — проверять отсортированы ли данные +
	-h — сортировать по числовому значению с учётом суффиксов

	Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	inputFile := flag.String("i", "", "")
	outputFile := flag.String("o", "", "")

	//column := flag.Bool("k", false, "")
	//numeric := flag.Int("n", 1, "")
	//reverse := flag.Bool("r", false, "")
	//unique := flag.Bool("u", false, "")

	//month := flag.Bool("M", false, "")
	//ignore := flag.Bool("b", false, "")
	//check := flag.Bool("c", false, "")
	//suffix := flag.Bool("h", false, "")

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

	sortedStrings := GetSortedStrings(data)

	err = WriteFileData(outputFilepath, sortedStrings)

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

func getNewline() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}

	return "\n"
}

func GetSortedStrings(data []string) []string {
	temporary := make([]string, len(data))

	copy(temporary, data)

	sort.Strings(temporary)

	return temporary
}

func GetReverseSortedStrings(data []string) []string {
	temporary := make([]string, len(data))

	copy(temporary, data)

	sort.Sort(sort.Reverse(sort.StringSlice(temporary)))

	return temporary
}

func GetSortedStringsByMonths(data []string) []string {
	temporary := make([]string, len(data))

	copy(temporary, data)

	months := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}

	sort.SliceStable(temporary, func(i, j int) bool {
		monthI := getMonthValue(temporary[i], months)
		monthJ := getMonthValue(temporary[j], months)

		return monthI < monthJ
	})

	return temporary
}

func getMonthValue(month string, months map[string]int) int {
	if value, isExist := months[strings.ToLower(month)]; isExist {
		return value
	}

	return 99
}

func GetSortedStringsWithKeyColumn(data []string, column int) []string {
	temporary := make([]string, len(data))

	copy(temporary, data)

	sort.SliceStable(temporary, func(i, j int) bool {
		columnsI := strings.Fields(temporary[i])
		columnsJ := strings.Fields(temporary[j])

		if column > 0 && column <= len(columnsI) && column <= len(columnsJ) {
			valueI := ""
			valueJ := ""

			if column-1 < len(columnsI) {
				valueI = columnsI[column-1]
			}

			if column-1 < len(columnsJ) {
				valueJ = columnsJ[column-1]
			}

			return valueI < valueJ
		}

		return temporary[i] < temporary[j]
	})

	return temporary
}

func GetSortedStringsByNumeric(data []string) []string {
	temporary := make([]string, len(data))

	copy(temporary, data)

	sort.SliceStable(temporary, func(i, j int) bool {
		numI, errI := strconv.Atoi(temporary[i])
		numJ, errJ := strconv.Atoi(temporary[j])

		if errI == nil && errJ == nil {
			return numI < numJ
		}

		return temporary[i] < temporary[j]
	})

	return temporary
}

func GetSortedStringsByNumericWithSuffix(data []string) []string {
	temporary := make([]string, len(data))

	copy(temporary, data)

	sort.SliceStable(temporary, func(i, j int) bool {
		valueI, suffixI := getNumericAndSuffix(temporary[i])
		valueJ, suffixJ := getNumericAndSuffix(temporary[j])

		if valueI < valueJ {
			return true
		}

		if valueI > valueJ {
			return false
		}

		return suffixI < suffixJ
	})

	return temporary
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
		dictionary[v] = struct{}{}
	}

	for v := range dictionary {
		temporary = append(temporary, v)
	}

	return temporary
}

func CheckSortedStrings(data []string) bool {
	temporary := make([]string, len(data))

	copy(temporary, data)

	sort.Strings(temporary)

	check := reflect.DeepEqual(data, temporary)

	return check
}

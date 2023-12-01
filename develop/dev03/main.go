package main

import (
	"bufio"
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

type FilesInterface interface {
	GetNewlineFile() string
	GetWorkDirectory() (string, error)
	GetFilePath(string) (string, error)
	GetFileData(string) ([]string, error)
	WriteFileData(string, []string) error
}

type Files struct {
	input  string
	output string
}

var _ FilesInterface = (*Files)(nil)

func (f *Files) GetFilePath(file string) (string, error) {
	wd, err := f.GetWorkDirectory()

	if err != nil {
		return "", err
	}

	filepath := path.Join(wd, file)

	return filepath, nil
}

func (f *Files) GetFileData(filepath string) ([]string, error) {
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

func (f *Files) WriteFileData(filepath string, data []string) error {
	file, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, v := range data {
		newline := f.GetNewlineFile()

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

func (f *Files) GetWorkDirectory() (string, error) {
	pwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return pwd, nil
}

func (f *Files) GetNewlineFile() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}

	return "\n"
}

func (f *Files) InitializeFiles() error {
	input := flag.Arg(0)

	if input == "" {
		return fmt.Errorf("укажите файл для чтения")
	}

	inputFilePath, err := f.GetFilePath(input)

	if err != nil {
		return err
	}

	f.input = inputFilePath

	output := flag.Arg(1)

	if output == "" {
		return fmt.Errorf("укажите файл для записи")
	}

	outputFilePath, err := f.GetFilePath(input)

	if err != nil {
		return err
	}

	f.output = outputFilePath

	return nil
}

type FlagsInterface interface {
	InitializeFlags()
}

type Flags struct {
	ColumnKey           int
	Numeric             bool
	Reverse             bool
	Unique              bool
	Month               bool
	IgnoreTrailingSpace bool
	Check               bool
	Suffix              bool
}

var _ FlagsInterface = (*Flags)(nil)

func (f *Flags) InitializeFlags() {
	columnKey := flag.Int("k", 0, "")
	numeric := flag.Bool("n", false, "")
	reverse := flag.Bool("r", false, "")
	unique := flag.Bool("u", false, "")

	month := flag.Bool("M", false, "")
	ignoreTrailingSpace := flag.Bool("b", false, "")
	check := flag.Bool("c", false, "")
	suffix := flag.Bool("h", false, "")

	flag.Parse()

	f.ColumnKey = *columnKey
	f.Numeric = *numeric
	f.Reverse = *reverse
	f.Unique = *unique

	f.Month = *month
	f.IgnoreTrailingSpace = *ignoreTrailingSpace
	f.Check = *check
	f.Suffix = *suffix
}

type ApplicationInterface interface {
	GetSortedStringsWithArguments([]string, *Flags) ([]string, error)
	GetStringsWithRemoveTrailingSpace([]string) []string
	GetUniqueStrings([]string) []string
	CheckSortedStrings([]string) bool
	GetNumericValue(string) int
	GetNumericAndSuffix(string) (int, string)
	GetSortColumnKey(string, int) string
	GetMonthValue(string) int
}

type Application struct{}

var _ ApplicationInterface = (*Application)(nil)

func main() {
	flags := &Flags{}
	files := &Files{}

	flags.InitializeFlags()

	err := files.InitializeFiles()

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	data, err := files.GetFileData(files.input)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	application := &Application{}

	result, err := application.GetSortedStringsWithArguments(
		data,
		flags,
	)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	err = files.WriteFileData(files.output, result)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}
}

func (a *Application) GetSortedStringsWithArguments(
	data []string,
	flags *Flags,
) ([]string, error) {
	temporary := make([]string, len(data))

	copy(temporary, data)

	if flags.Check && a.CheckSortedStrings(temporary) {
		return nil, fmt.Errorf("данные уже отсортированы")
	}

	if flags.IgnoreTrailingSpace {
		temporary = a.GetStringsWithRemoveTrailingSpace(temporary)
	}

	function := func(i, j int) bool {
		if flags.Month {
			return a.GetMonthValue(temporary[i]) < a.GetMonthValue(temporary[j])
		}

		if flags.ColumnKey > 0 {
			return a.GetSortColumnKey(temporary[i], flags.ColumnKey) < a.GetSortColumnKey(temporary[j], flags.ColumnKey)
		}

		return temporary[i] < temporary[j]
	}

	if flags.Numeric {
		function = func(i, j int) bool {
			return a.GetNumericValue(temporary[i]) < a.GetNumericValue(temporary[j])
		}
	}

	if flags.Suffix {
		function = func(i, j int) bool {
			valueI, suffixI := a.GetNumericAndSuffix(temporary[i])
			valueJ, suffixJ := a.GetNumericAndSuffix(temporary[j])

			if valueI < valueJ {
				return true
			}

			if valueI > valueJ {
				return false
			}

			return suffixI < suffixJ
		}
	}

	if flags.Unique {
		temporary = a.GetUniqueStrings(temporary)
	}

	sort.SliceStable(temporary, function)

	if flags.Reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(temporary)))
	}

	return temporary, nil
}

func (a *Application) GetStringsWithRemoveTrailingSpace(data []string) []string {
	temporary := make([]string, len(data))

	copy(temporary, data)

	for i, v := range temporary {
		temporary[i] = strings.TrimRightFunc(v, unicode.IsSpace)
	}

	return temporary
}

func (a *Application) GetUniqueStrings(data []string) []string {
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

func (a *Application) CheckSortedStrings(data []string) bool {
	check := sort.StringsAreSorted(data)

	return check
}

func (a *Application) GetNumericValue(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		return 999999999
	}

	return num
}

func (a *Application) GetNumericAndSuffix(input string) (int, string) {
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

func (a *Application) GetSortColumnKey(str string, column int) string {
	columns := strings.Fields(str)

	if column > 0 && column <= len(columns) {
		return columns[column-1]
	}

	return str
}

func (a *Application) GetMonthValue(month string) int {
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

	if _, isExist := months[strings.ToLower(month)]; !isExist {
		return 99
	}

	return months[strings.ToLower(month)]
}

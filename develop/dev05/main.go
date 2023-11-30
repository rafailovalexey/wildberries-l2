package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

/*
	=== Утилита grep ===

	Реализовать утилиту фильтрации (man grep)

	Поддержать флаги:
	-A - "after" печатать +N строк после совпадения
	-B - "before" печатать +N строк до совпадения
	-C - "context" (A+B) печатать ±N строк вокруг совпадения
	-c - "count" (количество строк)
	-i - "ignore-case" (игнорировать регистр)
	-v - "invert" (вместо совпадения, исключать)
	-F - "fixed", точное совпадение со строкой, не паттерн
	-n - "line num", печатать номер строки

	Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type FilesInterface interface {
	GetFilePath(string) (string, error)
	GetFileData(string) ([]string, error)
	GetWorkDirectory() (string, error)
}

type Files struct {
	input string
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

func (f *Files) GetWorkDirectory() (string, error) {
	pwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return pwd, nil
}

func (f *Files) InitializeFiles() error {
	input := flag.Arg(0)

	if input == "" {
		return errors.New("укажите файл для чтения")
	}

	filepath, err := f.GetFilePath(input)

	if err != nil {
		return err
	}

	f.input = filepath

	return nil
}

type FlagsInterface interface {
	InitializeFlags()
}

type Flags struct {
	After       int
	Before      int
	Context     int
	Count       bool
	IgnoreCase  bool
	Invert      bool
	Fixed       bool
	LineNumbers bool
}

var _ FlagsInterface = (*Flags)(nil)

func (f *Flags) InitializeFlags() {
	after := flag.Int("A", 0, "")
	before := flag.Int("B", 0, "")
	context := flag.Int("C", 0, "")
	count := flag.Bool("c", false, "")

	ignoreCase := flag.Bool("i", false, "")
	invert := flag.Bool("v", false, "")
	fixed := flag.Bool("F", false, "")
	lineNumbers := flag.Bool("n", false, "")

	flag.Parse()

	f.After = *after
	f.Before = *before
	f.Context = *context
	f.Count = *count
	f.IgnoreCase = *ignoreCase
	f.Invert = *invert
	f.Fixed = *fixed
	f.LineNumbers = *lineNumbers
}

type ApplicationInterface interface {
	Grep([]string, string, *Flags) ([]string, error)
	Match(string, string, *Flags) (bool, error)
	PrintFileData([]string, *Flags)
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

	pattern := flag.Arg(1)

	application := &Application{}

	data, err = application.Grep(data, pattern, flags)

	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	application.PrintFileData(
		data,
		flags,
	)
}

func (a *Application) Grep(data []string, pattern string, flags *Flags) ([]string, error) {
	var temporary []string

	for index, line := range data {
		matched, err := a.Match(line, pattern, flags)

		if err != nil {
			return nil, err
		}

		if (flags.Before > 0 && index < flags.Before) || (flags.After > 0 && index+flags.After >= len(data)) {
			temporary = append(temporary, line)
		} else if matched {
			temporary = append(temporary, line)
		} else if flags.Context > 0 && index+flags.Context < len(data) {
			temporary = append(temporary, data[index:index+flags.Context+1]...)
		}
	}

	return temporary, nil
}

func (a *Application) Match(line string, pattern string, flags *Flags) (bool, error) {
	if flags.Fixed {
		return line == pattern, nil
	}

	if flags.IgnoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}

	matched, err := regexp.MatchString(pattern, line)

	if err != nil {
		return false, err
	}

	return matched != flags.Invert, nil
}

func (a *Application) PrintFileData(
	data []string,
	flags *Flags,
) {
	for i, v := range data {
		if flags.Count {
			fmt.Printf("%d\n", len(data))

			break
		}

		if flags.LineNumbers {
			fmt.Printf("%d:%s\n", i+1, v)
		} else {
			fmt.Printf("%s\n", v)
		}
	}
}

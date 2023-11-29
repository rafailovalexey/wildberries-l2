package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
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
	//key := flag.String("k", "", "")
	//numeric := flag.String("n", "", "")
	//reverse := flag.String("r", "", "")
	//unique := flag.String("u", "", "")
	//
	//month := flag.String("M", "", "")
	//ignore := flag.String("b", "", "")
	//check := flag.String("c", "", "")
	//suffix := flag.String("h", "", "")

	flag.Parse()

	pwd, _ := os.Getwd()

	file := flag.Args()[0]
	filepath := path.Join(pwd, file)

	fmt.Println(filepath)

	openFile, err := os.OpenFile(filepath, os.O_RDWR, 0755)

	if err != nil {
		log.Fatalf("%v", err)

		return
	}

	defer openFile.Close()

	data := make([]byte, 0, 100)

	n, err := openFile.Read(data)

	fmt.Println(n)

	if err != nil {
		log.Fatalf("%v", err)

		return
	}

	fmt.Println(data)
}

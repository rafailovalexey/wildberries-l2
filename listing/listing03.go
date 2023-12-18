package main

import (
	"log"
	"os"
)

func Foo() error {
	var err *os.PathError = nil

	return err
}

func main() {
	err := Foo()

	log.Printf("%v\n", err)
	log.Printf("%v\n", err == nil)
}

package main

import (
	"log"
)

func test() (x int) {
	defer func() {
		x++
	}()

	x = 1

	return
}

func anotherTest() int {
	var x int

	defer func() {
		x++
	}()

	x = 1

	return x
}

func main() {
	log.Printf("%d\n", test())
	log.Printf("%d\n", anotherTest())
}

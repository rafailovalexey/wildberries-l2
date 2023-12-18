package main

import (
	"log"
)

func main() {
	var s = []string{"1", "2", "3"}

	modifySlice(s)

	log.Printf("%s\n", s)
}

func modifySlice(i []string) {
	i[0] = "3"

	i = append(i, "4")

	i[1] = "5"

	i = append(i, "6")
}

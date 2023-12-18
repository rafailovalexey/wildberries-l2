package main

import (
	"log"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}

	var b []int = a[1:4]

	log.Printf("%d\n", b)
}

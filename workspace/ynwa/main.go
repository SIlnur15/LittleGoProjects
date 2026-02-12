package main

import (
	"fmt"
	"unicode"
)

func main() {
	var (
		input  string
		export int
	)
	fmt.Scan(&input)
	for _, elem := range input {
		if unicode.IsDigit(elem) {
			export += 1
		}
	}
	fmt.Printf("Количество цифр в строке: %d\n", export)
}

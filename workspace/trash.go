package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	for _, letter := range input {
		if strings.Count(input, string(letter)) == 2 {
			fmt.Println(string(letter))
			break
		}
	}
}

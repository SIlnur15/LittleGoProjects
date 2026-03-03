package main

import (
	"fmt"
)

func main() {
	word := "hello"
	arr := []byte(word)
	for _, elem := range arr {
		fmt.Println(elem)
	}
}

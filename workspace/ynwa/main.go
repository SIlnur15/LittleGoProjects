package main

import "fmt"

func main() {
	var second string
	second = `defgs`
	first := []rune(second)
	fmt.Println(len(first)) // 5
}

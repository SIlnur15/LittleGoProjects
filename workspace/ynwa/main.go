package main

import (
	"fmt"
)

func isPalindrome(input string) bool {
	arr := []rune(input)
	exp := make([]rune, len(input))
	output := ""
	for i, letter := range arr {
		exp[len(arr)-i-1] = letter
	}
	for _, runedLetter := range exp {
		output += string(runedLetter)
	}
	if input == output {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome("dected"))
}

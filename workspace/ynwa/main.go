package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func mapper(w rune) rune {
	switch w {
	case '!', '-', ',', '?', '.', ';', ':':
		return 0 // далее в теле программы нужно будет избавляться от пустых рун
		// потому что, как оказалось, их хоть и не видно, но они есть
	}
	return w
}

func main() {
	// Считывание всей строки
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()

	export := strings.Map(mapper, str)
	exportRunes := []rune(export)
	sample := []rune{}
	for idx, elem := range exportRunes {
		if elem != 0 {
			sample = append(sample, exportRunes[idx])
		}
	}
	word := string(sample)        // получаю строку из среза рун
	words := strings.Fields(word) // Разбиение строки на слова по пробелам

	// Поиск самого длинного слова
	longestWord := ""
	for _, word := range words {
		if utf8.RuneCountInString(word) > utf8.RuneCountInString(longestWord) {
			longestWord = word
		}
	}
	fmt.Println(longestWord)
}

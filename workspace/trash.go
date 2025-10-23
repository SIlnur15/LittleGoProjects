package main

import (
    "fmt"
    "sort"
)

// Функция для проверки, являются ли две строки анаграммами
func isAnagram(s1, s2 string) bool {
    // Если длины строк не совпадают, они не могут быть анаграммами
    if len(s1) != len(s2) {
        return false
    }
    
    // Преобразуем строки в срезы рун (для корректной работы с Unicode)
    r1 := []rune(s1)
    r2 := []rune(s2)
    
    // Сортируем оба среза
    sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
    sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })
    
    // Сравниваем отсортированные срезы
    return string(r1) == string(r2)
}

func main() {
	// не изменяйте функцию main!!

	var word1, word2 string
	fmt.Scan(&word1, &word2)

	if isAnagram(word1, word2) {
		fmt.Println("Да")
	} else {
		fmt.Println("Нет")
	}
}
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int) // карта для хранения индексов каждого символа
	maxLen := 0
	start := 0 // начальный индекс текущей подстроки без дубликатов
	count := 1

	for i, ch := range s {
		fmt.Println(count)
		fmt.Println(charIndex)
		if lastIdx, ok := charIndex[ch]; ok && lastIdx >= start {
			// если символ уже есть в мапе
			fmt.Println(lastIdx, ok)
			// если символ повторяется, то перемещаемся на следующую позицию после последнего появления
			start = lastIdx + 1
		}
		charIndex[ch] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		fmt.Println(charIndex)
		count += 1
		fmt.Println(maxLen)
		fmt.Println("========")
	}

	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s)) // Вывод: 3
}

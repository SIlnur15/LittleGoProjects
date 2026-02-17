package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Создаем новый сканер из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)

	// Устанавливаем разделитель
	scanner.Split(bufio.ScanWords)

	// Читаем слова из ввода и выводим их на экран
	for scanner.Scan() {
		word := scanner.Text()
		fmt.Println(word)
	}
}

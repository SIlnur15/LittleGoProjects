package main

import "fmt"

// Callback-функция как параметр
func processNumbers(numbers []int, callback func(int)) {
	for _, num := range numbers {
		callback(num) // Вызов callback для каждого числа
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Передаем анонимную функцию как callback
	processNumbers(numbers, func(n int) {
		fmt.Printf("Число: %d\n", n)
	})
}

package main

import (
	"fmt"
	"sync"
)

func processBatch(data []int) []int {
	var wg sync.WaitGroup
	results := make([]int, len(data))

	for i, item := range data {
		i, item := i, item // Создаём локальные копии
		wg.Go(func() {
			// Параллельная обработка
			results[i] = item * 2
		})
	}

	wg.Wait()
	return results
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	result := processBatch(data)
	fmt.Println(result) // [2 4 6 8 10]
}

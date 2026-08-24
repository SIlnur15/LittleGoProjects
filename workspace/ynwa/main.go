package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// generateTriangular — это функция-генератор (Producer).
// ЕЕ ОТВЕТСТВЕННОСТЬ:
// 1. Генерировать треугольные числа.
// 2. Писать их в канал numbers.
// 3. Закрыть канал numbers, когда числа закончились.
// 4. Внимательно слушать ctx.Done() (сигнал quit), чтобы быстро завершиться, если попросят.
func generateTriangular(ctx context.Context, numbers chan<- int, limit int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(numbers) // Правило: продюсер закрывает канал данных

	for i := 1; i <= limit; i++ {
		// Вычисляем треугольное число
		triNum := i * (i + 1) / 2

		select {
		case numbers <- triNum:
			// Успешно отправили число
		case <-ctx.Done():
			// Получен сигнал отмены (quit) от main!
			// Генератор не закрывает quit, он только реагирует на него.
			return
		}

		time.Sleep(100 * time.Millisecond) // Имитация полезной нагрузки для наглядности
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	numbers := make(chan int)
	wg.Add(1)
	go generateTriangular(ctx, numbers, 50, &wg)
	for num := range numbers {
		if num < 50 {
			fmt.Println(num)
		} else {
			fmt.Println(num)
			cancel()
			break
		}
	}
	wg.Wait()
	fmt.Println("Прекращаем генерацию чисел")
}

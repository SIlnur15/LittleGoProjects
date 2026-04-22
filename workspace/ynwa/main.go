package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	fmt.Printf("Горутина %d: запущена\n", id)

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Горутина %d: отмена! Причина: %v\n", id, ctx.Err())
			// Здесь можно добавить очистку ресурсов
			return
		default:
			fmt.Printf("Горутина %d: итерация %d\n", id, i)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Запускаем несколько горутин
	for i := 1; i <= 2; i++ {
		go worker(ctx, i)
	}

	// Ждём завершения всех горутин
	time.Sleep(5 * time.Second)
	fmt.Println("Основная программа завершена")
}

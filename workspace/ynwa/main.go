package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) error {
	// Имитация работы, которая может быть прервана по таймауту
	for i := 0; i < 5; i++ {
		// Проверяем, не истёк ли таймаут или не отменён ли контекст
		if ctx.Err() != nil {
			return ctx.Err() // возвращаем ошибку контекста (таймаут или отмена)
		}
		fmt.Println("Работаем...", i)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := doWork(ctx)
	if err != nil {
		fmt.Println("Операция прервана:", err)
	} else {
		fmt.Println("Операция завершена успешно")
	}
}

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Создаем базовый пустой контекст (Background)
	// и оборачиваем его в контекст с таймаутом на 2 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	// ОБЯЗАТЕЛЬНО вызываем cancel для освобождения ресурсов
	// по завершении работы функции (даже если таймаут не наступил)
	defer cancel()

	// Запускаем долгую операцию
	go slowOperation(ctx)

	// Ждем, пока контекст завершится (по таймауту или через cancel)
	<-ctx.Done()

	// Проверяем, почему именно завершился контекст
	fmt.Println("Основной поток: Контекст завершен по причине:", ctx.Err())
}

func slowOperation(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // Имитируем работу на 5 секунд
		fmt.Println("Операция успешно завершена!")
	case <-ctx.Done(): // Сработает через 2 секунды, так как контекст закроется
		fmt.Println("slowOperation: Время вышло, прекращаю работу...")
	}
}

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Создаём контекст с дедлайном через 5 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Получаем дедлайн и флаг
	deadline, ok := ctx.Deadline()

	if ok {
		fmt.Printf("Дедлайн установлен: %v\n", deadline)
		fmt.Printf("Осталось времени: %v\n", time.Until(deadline))
	} else {
		fmt.Println("Дедлайн не установлен")
	}
}

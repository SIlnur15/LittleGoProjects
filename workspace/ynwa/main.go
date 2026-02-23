package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем объект времени
	t := time.Date(2023, time.October, 17, 10, 0, 0, 0, time.UTC)

	// Маршалинг в JSON
	jsonBytes, err := t.MarshalJSON()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Вывод результата
	fmt.Printf("JSON-представление: %s\n", string(jsonBytes))
}

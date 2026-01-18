package main

import (
	"errors"
	"fmt"
)

func main() {
	err := doSomething()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err) // Выведет обёрнутую ошибку

		// Проверяем, содержится ли исходная ошибка в цепочке
		if errors.Is(err, ErrOriginal) {
			fmt.Println("Да, это исходная ошибка!")
		}
	}
}

var ErrOriginal = errors.New("исходная ошибка")

func doSomething() error {
	// Имитируем ошибку
	err := ErrOriginal

	// Оборачиваем ошибку с дополнительным контекстом
	return fmt.Errorf("не удалось выполнить действие: %w", err)
}

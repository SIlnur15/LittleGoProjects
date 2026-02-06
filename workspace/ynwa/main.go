package main

import "fmt"

// Функция, которая может вызвать панику
func riskyFunction() {
	panic("что-то пошло не так!")
}

// Функция для восстановления после паники
func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Паника перехвачена:", r)
	}
}

func main() {
	// Откладываем вызов обработчика до завершения main
	defer handlePanic()

	fmt.Println("Начало работы")
	riskyFunction() // вызывает panic
	fmt.Println("Эта строка не будет напечатана")
}

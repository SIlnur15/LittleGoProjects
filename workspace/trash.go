package main

import (
	"fmt"
	"os"
)

func main() {
	// Создаём файл (если он существует, он будет перезаписан)
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close() // Закрываем файл после завершения работы

	// Текст для записи
	text := "Привет, мир! Это текст, записанный в файл.\n"

	// Записываем текст в файл
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	fmt.Println("Данные успешно записаны в файл.")
}

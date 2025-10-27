package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// открываем CSV файл
	file, err := os.OpenFile("output.csv", os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	data := [][]string{ // данные для вставки
		{"Василий", "68", "Вена"},
		{"Виталина", "120", "Таллин"},
	}
	for _, record := range data {
		if err := writer.Write(record); err != nil {
			fmt.Println("Ошибка записи строки:", err)
			return
		}
	}
}

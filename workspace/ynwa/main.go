package main

import (
	"fmt"
	"os"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Ошибка при получении домашней директории:", err)
		return
	}
	fmt.Println("Домашняя директория пользователя:", homeDir)
}

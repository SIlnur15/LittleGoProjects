package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Основной логгер — в stderr
	mainLog := log.New(os.Stderr, "MAIN: ", log.LstdFlags)

	// Логгер для ошибок — в stderr с другим префиксом
	errLog := log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)

	// Логгер в файл
	file, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		// Используем errLog.Fatalf — он выведет красивую ошибку с префиксом ERROR и завершит программу
		errLog.Fatalf("Не удалось открыть файл debug.log: %v", err)
	}
	defer file.Close()

	debugLog := log.New(io.MultiWriter(file, os.Stdout), "DEBUG: ", log.LstdFlags)
	mainLog.Println("Приложение запущено")
	errLog.Println("Невозможно подключиться к устройству")
	debugLog.Println("Полный дамп данных...")
}

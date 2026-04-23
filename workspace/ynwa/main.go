package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// Создаём HTTP-клиент с таймаутом
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Создаём запрос
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

	// Создаём контекст с таймаутом 2 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Привязываем контекст к запросу
	req = req.WithContext(ctx)

	// Выполняем запрос
	resp, err := client.Do(req)
	if err != nil {
		// Проверяем, была ли ошибка из-за отмены контекста
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Запрос отменён по таймауту контекста")
		} else {
			fmt.Println("Ошибка выполнения запроса:", err)
		}
		return
	}
	defer resp.Body.Close()

	// Читаем тело ответа (с проверкой контекста)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Чтение ответа прервано по таймауту контекста")
		} else {
			fmt.Println("Ошибка чтения ответа:", err)
		}
		return
	}

	// Выводим результат
	fmt.Printf("Получено %d байт. Статус: %s\n", len(data), resp.Status)
}

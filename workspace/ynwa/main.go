package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Список URL источников данных для запросов
	urls := []string{
		"https://source1.com/api",
		"https://source2.com/api",
		"https://source3.com/api",
	}

	// Создаем контекст с возможностью отмены
	// Это позволит прервать все pending-запросы после получения первого ответа
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Гарантируем освобождение ресурсов при выходе

	// Буферизированный канал для получения результатов
	// Размер буфера = количеству URL, чтобы избежать блокировок
	result := make(chan string, len(urls))

	// Запускаем горутины для параллельных запросов ко всем источникам
	for _, url := range urls {
		go func(u string) {
			// Создаем HTTP-запрос с привязкой к контексту
			// Контекст автоматически отменит запрос при вызове cancel()
			req, _ := http.NewRequestWithContext(ctx, "GET", u, nil)

			// Выполняем HTTP-запрос
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return // Игнорируем ошибки соединения
			}
			defer resp.Body.Close() // Гарантируем закрытие тела ответа

			// Проверяем успешный статус ответа (200 OK)
			if resp.StatusCode == http.StatusOK {
				// Отправляем результат в канал
				result <- fmt.Sprintf("Ответ от %s", u)
				// Отменяем контекст - прерываем все остальные запросы
				cancel()
			}
		}(url)
	}

	// Ожидаем первый успешный ответ или таймаут
	select {
	case res := <-result:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Таймаут")
	}
}

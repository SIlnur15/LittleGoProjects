package main

import (
	"fmt"
	"net"
	"time"
)

// checkPort пытается установить TCP-соединение с localhost на указанный порт.
// Результат проверки (открыт или закрыт порт) отправляется в канал results.
func checkPort(port int, results chan<- string) {
	address := fmt.Sprintf("localhost:%d", port) // Формируем адрес в формате "localhost:порт"

	// Пытаемся установить соединение с таймаутом 2 секунды
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err == nil {
		// Если соединение успешно, закрываем его и отправляем сообщение что порт открыт
		conn.Close()
		results <- fmt.Sprintf("Port %d is OPEN", port)
	} else {
		// Если ошибка — порт закрыт или недоступен
		results <- fmt.Sprintf("Port %d is CLOSED", port)
	}
}

func main() {
	// Список портов, которые будем проверять
	ports := []int{80, 443, 8080, 22, 3306}

	// Создаём канал для передачи результатов проверки
	results := make(chan string)

	// Запускаем горутину для проверки каждого порта
	for _, port := range ports {
		go checkPort(port, results)
	}

	// Получаем результаты из канала и выводим их на экран
	// Цикл выполняется ровно столько раз, сколько портов для проверки
	for range ports {
		fmt.Println(<-results)
	}
}

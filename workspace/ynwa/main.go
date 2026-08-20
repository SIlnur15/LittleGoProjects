package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// worker - функция-воркер для сканирования портов
func worker(ports <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик при завершении работы воркера

	for port := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			conn.Close()
			results <- fmt.Sprintf("Port %d is open", port)
		} else {
			results <- fmt.Sprintf("Port %d is closed", port)
		}
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan string)
	var wg sync.WaitGroup

	// Запускаем пул воркеров
	for i := 0; i < cap(ports); i++ {
		wg.Add(1)
		go worker(ports, results, &wg)
	}

	// Горутина для сбора результатов
	go func() {
		wg.Wait()
		close(results)
	}()

	// Отправляем номера портов для сканирования
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
		close(ports) // Закрываем канал портов после отправки всех заданий
	}()

	// Выводим результаты
	for result := range results {
		fmt.Println(result)
	}
}

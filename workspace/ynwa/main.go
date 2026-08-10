package main

import (
	"fmt"
	"math/rand"
	"time"
)

// simulateWork имитирует работу сервиса (разное время выполнения)
func simulateWork(id int, resultChan chan<- string) {
	delay := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(delay)
	resultChan <- fmt.Sprintf("Результат от сервиса %d (задержка: %v)", id, delay)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	resultChan := make(chan string, 3) // Буферизованный канал для результатов

	// Запускаем несколько горутин для обработки запроса
	go simulateWork(1, resultChan)
	go simulateWork(2, resultChan)
	go simulateWork(3, resultChan)

	// Берем первый ответ и завершаем работу
	firstResult := <-resultChan
	fmt.Println("Первый ответ:", firstResult)
}

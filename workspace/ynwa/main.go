package main

import (
	"fmt"
	"time"
)

func worker(chChan chan chan string) {
	replyChan := make(chan string) // 1. Создаем свой персональный канал для ответа
	chChan <- replyChan            // 2. Отправляем НАШ канал в канал каналов
	msg := <-replyChan             // 3. Ждем, пока через наш личный канал нам пришлют ответ
	fmt.Println("Воркер получил сообщение:", msg)
}

func main() {
	chChan := make(chan chan string)        // Создаем канал, который может принимать только другие каналы строк
	go worker(chChan)                       // Запускаем воркера
	personalWorkerChan := <-chChan          // Ждем, пока воркер пришлет нам ссылку на свой персональный канал
	personalWorkerChan <- "Привет из main!" // Теперь мы пишем данные напрямую в канал этого воркера!
	time.Sleep(time.Millisecond * 100)      // Даем время воркеру на печать
}

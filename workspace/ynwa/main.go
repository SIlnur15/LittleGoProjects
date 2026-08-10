package main

import (
	"log"
	"time"
)

func worker(heartbeat chan<- struct{}, done <-chan struct{}) {
	defer close(heartbeat)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			heartbeat <- struct{}{} // Отправляем пульс
		case <-time.After(2 * time.Second):
			// Долгая операция (может быть обёрнута в отдельный select)
			log.Println("working...")
		}
	}
}

func main() {
	done := make(chan struct{})
	heartbeat := make(chan struct{})
	go worker(heartbeat, done)

	// Проверка "пульса" горутины
	select {
	case <-heartbeat:
		log.Println("Горутина жива")
	case <-time.After(1 * time.Second):
		log.Fatal("Горутина не отвечает!")
	}
}

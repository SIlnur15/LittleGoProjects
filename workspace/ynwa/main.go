package main

import (
	"fmt"
	"time"
)

func main() {
	synchronization()
}

func synchronization() {
	ch := make(chan struct{}) // Канал для синхронизации

	go func() {
		fmt.Println("Горутина начинает работу")
		time.Sleep(time.Second)
		fmt.Println("Горутина завершила работу")
		ch <- struct{}{} // Сигнал о завершении
	}()

	<-ch // Ожидание завершения
	fmt.Println("Основная горутина продолжает работу")
}

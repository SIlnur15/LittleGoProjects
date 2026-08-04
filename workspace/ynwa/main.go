package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	taskChannel := make(chan string) // Создаем небуферизированный канал для передачи задач

	wg.Add(1)
	go func() { // Горутина-исполнитель
		defer wg.Done()
		for task := range taskChannel {
			// Получение задачи (блокируется, пока не будет отправки)
			fmt.Printf("start of processing: %s\n", task)
			time.Sleep(1 * time.Second) // Имитация работы
			fmt.Printf("is over: %s\n", task)
		}
	}()

	tasks := []string{"task 1", "task 2", "task 3"}
	for _, task := range tasks { // Отправляем 3 задачи последовательно
		fmt.Printf("sending %s...\n", task)
		taskChannel <- task
		fmt.Printf("was sent: %s\n", task)
	}

	close(taskChannel) // Закрываем канал после отправки всех задач
	wg.Wait()
}

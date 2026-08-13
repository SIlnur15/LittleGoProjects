package main

import (
	"fmt"
	"sync"
	"time"
)

func workerWithWG(wg *sync.WaitGroup, stop <-chan struct{}, id int) {
	defer wg.Done()

	for {
		select {
		case <-stop:
			fmt.Printf("Worker %d received stop signal\n", id)
			return
		default:
			fmt.Printf("Worker %d working...\n", id)
			time.Sleep(time.Duration(id) * 500 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	stop := make(chan struct{})

	// Запускаем горутины
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerWithWG(&wg, stop, i)
	}

	// Даем поработать 5 секунд
	time.Sleep(5 * time.Second)

	// Останавливаем все горутины
	close(stop)

	// Ждем завершения всех горутин
	wg.Wait()
	fmt.Println("All workers stopped")
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	manyWritersOneReader()
}

func manyWritersOneReader() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ch <- id
		}(i)
	}

	// Отдельная горутина для ожидания и закрытия канала
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Читатель успешно завершит цикл, когда канал закроется
	for id := range ch {
		fmt.Println("Получен ID:", id)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data    string
	rwMutex sync.RWMutex
)

func writer(id int) {
	rwMutex.Lock() // Закрыл дверь писателей
	defer rwMutex.Unlock()

	fmt.Printf("Писатель %d вошел\n", id)
	data = fmt.Sprintf("данные от писателя %d", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Писатель %d вышел\n", id)
}

func reader(id int) {
	rwMutex.RLock() // Проверил, нет ли писателя
	defer rwMutex.RUnlock()

	fmt.Printf("Читатель %d вошел\n", id)
	fmt.Printf("Читатель %d прочитал: %s\n", id, data)
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("Читатель %d вышел\n", id)
}

func main() {
	// Запускаем писателей и читателей
	for i := 0; i < 3; i++ {
		go writer(i)
	}

	for i := 0; i < 5; i++ {
		go reader(i)
	}

	time.Sleep(2 * time.Second)
}

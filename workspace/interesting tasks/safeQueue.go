package main

import (
	"fmt"
	"sync"
	"time"
)

// структура-очередь с условной переменной
type Queue struct {
	mu    sync.Mutex
	cond  *sync.Cond
	items []int
}

// Конструктор
func NewQueue() *Queue {
	q := &Queue{}
	q.cond = sync.NewCond(&q.mu)
	return q
}

// Добавить элемент
func (q *Queue) Push(item int) {
	q.mu.Lock()
	q.items = append(q.items, item)
	fmt.Printf("Добавлено: %d\n", item)
	q.cond.Signal() // Уведомить одного ожидающего
	q.mu.Unlock()
}

// Взять элемент (блокируется, если пусто)
func (q *Queue) Pop() int {
	q.mu.Lock() // Блокировка
	// Ждём, пока очередь не станет непустой
	for len(q.items) == 0 { // при работе с Cond только for, а не if !
		q.cond.Wait() // Разблокирует mu, ждёт сигнала, потом снова блокирует
	}
	// Берём первый элемент
	item := q.items[0]
	q.items = q.items[1:]
	q.mu.Unlock() // Разблокировка
	return item
}

func main() {
	var n int
	fmt.Scan(&n)

	q := NewQueue()

	// Потребитель — ждёт элементы и выводит
	go func() {
		for i := 0; i < n; i++ {
			val := q.Pop()
			fmt.Printf("Получено: %d\n", val)
			time.Sleep(100 * time.Millisecond) // Имитация обработки
		}
	}()

	// Производитель — добавляет элементы
	for i := 1; i <= n; i++ {
		q.Push(i)
		time.Sleep(30 * time.Millisecond) // Добавляем  чаще
	}

	// Ждём, чтобы потребитель успел всё обработать
	time.Sleep(2 * time.Second)
	fmt.Println("Готово!")
}

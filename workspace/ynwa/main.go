package main

import (
	"fmt"
	"time"
)

// producer принимает канал только для записи (chan<-)
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i // отправляем значения в канал
		fmt.Printf("Отправлено: %d\n", i)
		time.Sleep(time.Second) // задержка для наглядности
	}
	close(ch) // закрываем канал после отправки всех значений
}

// consumer принимает канал только для чтения (<-chan)
func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Получено: %d\n", num)
	}
}

func main() {
	ch := make(chan int) // создаем двунаправленный канал

	go producer(ch) // запускаем producer в горутине
	consumer(ch)    // consumer работает в main горутине

	fmt.Println("Готово!")
}

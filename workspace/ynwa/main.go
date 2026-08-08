package main

import (
	"fmt"
	"time"
)

func tee(input <-chan int, outputs ...chan<- int) {
	for v := range input {
		for _, output := range outputs {
			output <- v
		}
	}

	// Закрываем все выходные каналы после завершения
	for _, output := range outputs {
		close(output)
	}
}

func main() {
	in := make(chan int)
	out1 := make(chan int)
	out2 := make(chan int)

	// Запускаем tee в отдельной горутине
	go tee(in, out1, out2)

	// Запускаем потребителей
	go func() {
		for v := range out1 {
			fmt.Println("out1:", v)
		}
	}()

	go func() {
		for v := range out2 {
			fmt.Println("out2:", v)
		}
	}()

	// Отправляем данные в входной канал
	for i := 0; i < 5; i++ {
		in <- i
	}

	close(in)
	time.Sleep(time.Second) // Даем время для обработки
}

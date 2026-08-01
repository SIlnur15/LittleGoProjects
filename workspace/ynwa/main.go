package main

import (
	"fmt"
)

func sayHello(ch chan string) {
	ch <- "Hello to chanel!"
}

func main() {
	ch := make(chan string) // инициализируем канал через make

	go sayHello(ch)   // Запуск горутины, которая отправит сообщение
	fmt.Println(<-ch) // Получение и вывод сообщения из канала
	fmt.Println("the main stream is over!")
}

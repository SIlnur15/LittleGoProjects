package main

import "fmt"

func main() {
	ch1 := make(chan int) // небуферизованный канал
	ch2 := make(chan int) // небуферизованный канал

	go func() {
		ch1 <- 1 // отправка в ch1
		fmt.Println("Sent 1 to ch1")
		<-ch2 // получение из ch2
	}()

	go func() {
		ch2 <- 2 // отправка в ch2
		fmt.Println("Sent 2 to ch2")
		<-ch1 // получение из ch1
	}()

	select {} // бесконечное ожидание
}

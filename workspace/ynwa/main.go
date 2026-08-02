package main

import "fmt"

func routiner(cch chan int) {
	cch <- 15
	cch <- 25
}

func main() {
	channel := make(chan int)
	go routiner(channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

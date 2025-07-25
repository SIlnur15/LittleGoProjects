package main

import "fmt"

func toPrint(mc chan int) {
	mc <- 555
}

func main() {
	var myChanel chan int
	myChanel = make(chan int)
	go toPrint(myChanel)
	fmt.Println(<-myChanel)
}

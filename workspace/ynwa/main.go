package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("hello")
}

func main() {
	go sayHello()
	time.Sleep(1 * time.Millisecond)
}

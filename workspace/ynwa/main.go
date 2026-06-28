package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func main() {
	go sayHello("Haidaric")
	fmt.Println("hello hello")
	time.Sleep(1 * time.Second)
}

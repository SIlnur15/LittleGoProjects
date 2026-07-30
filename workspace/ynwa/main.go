package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("hello GOOOOOOOO") // anonymus goroutine
	}()
	fmt.Println("hello Go")
	time.Sleep(time.Millisecond)
}

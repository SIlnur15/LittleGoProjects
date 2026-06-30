package main

import (
	"fmt"
	"time"
)

func main() {
	var x, y int
	go func() {
		x = 1
		fmt.Print("у:", y, " ")
	}()
	go func() {
		y = 1
		fmt.Print("х:", x, " ")
	}()
	time.Sleep(1)
}

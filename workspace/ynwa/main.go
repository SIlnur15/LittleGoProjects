package main

import (
	"fmt"
	"time"
)

func main() {
	d := 5*time.Minute + 10*time.Hour // 10 hours 5 минут
	fmt.Println(d)                    // 10h5m0s
}

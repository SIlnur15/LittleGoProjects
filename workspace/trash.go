package main

import (
	"fmt"
)

func main() {
	i := 2
	var n int
	fmt.Scan(&n)
	for n%i != 0 {
		i++
	}
	fmt.Println(i)
}

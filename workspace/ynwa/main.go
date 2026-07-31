package main

import (
	"fmt"
)

func change(x int) int {
	x = 22
	return x
}

func main() {
	x := 5
	fmt.Println(x)
	a := change(x)
	fmt.Println(a)
	fmt.Println(x)
}

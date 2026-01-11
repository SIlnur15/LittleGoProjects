package main

import (
	"fmt"
)

func main() {
	var a interface{}
	a = 15
	switch a.(type) {
	case int:
		fmt.Println("that's int") // that's int
	default:
		fmt.Println("kek")
	}
}

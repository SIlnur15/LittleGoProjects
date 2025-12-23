package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	m["key"] = 15
	fmt.Println(m) // panic
}

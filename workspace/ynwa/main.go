package main

import (
	"fmt"
)

func main() {
	s := "Go123"
	for _, v := range s {
		fmt.Print(v, " ")
	}
}

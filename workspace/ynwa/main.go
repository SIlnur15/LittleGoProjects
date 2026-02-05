package main

import (
	"errors"
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if a == 5 {
		b := errors.New("nothing was entered")
		fmt.Println(b)
	} else {
		fmt.Println(a)
	}
}

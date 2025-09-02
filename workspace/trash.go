package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Print(1)

	if x != 5 {
		fmt.Print(2)
	}
	fmt.Print(3)
}

package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	fmt.Printf("%T\n", input/2)
	fmt.Println(input / 2)
}

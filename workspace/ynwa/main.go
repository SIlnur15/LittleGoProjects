package main

import "fmt"

type Person struct {
	string
	int
}

func main() {
	p1 := Person{"Alice", 30}
	fmt.Println(p1) // {Alice 30}
}

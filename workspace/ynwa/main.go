package main

import "fmt"

type Contact struct {
	name string
	age  int
}

func main() {
	x := Contact{"Klopp", 15}
	fmt.Printf("%#v\n", x)
}

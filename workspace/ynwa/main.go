package main

import "fmt"

func main() {
	a := func(name string) {
		fmt.Println("hello", name)
	}
	a("Roma")
	a("Sasha")
}

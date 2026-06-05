package main

import "fmt"

type myType struct {
	name string
	age  int
}

func (m myType) get() {
	fmt.Printf("hello, I'm %s\n", m.name)
}

func main() {
	a := myType{name: "Axel", age: 25}
	a.get()
}

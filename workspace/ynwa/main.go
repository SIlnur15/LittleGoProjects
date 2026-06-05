package main

import "fmt"

type myType struct {
	name string
	age  int
}

func (m myType) get() {
	fmt.Printf("hello, I'm %s, I'm %d\n", m.name, m.age)
}

func (m *myType) set() {
	m.age = m.age + 5
}

func main() {
	a := myType{name: "Axel", age: 25}
	a.get()
	a.set()
	a.get()
}

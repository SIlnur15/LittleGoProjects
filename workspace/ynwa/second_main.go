package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	animals := []Animal{Dog{Name: "Fido"}, Cat{Name: "Fluffy"}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

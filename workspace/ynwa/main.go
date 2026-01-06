package main

import (
	"fmt"
)

func getPerson(n string, a int) struct {
	name string
	age  int
} {
	return struct {
		name string
		age  int
	}{
		name: n,
		age:  a,
	}
}

func main() {
	person := getPerson("Miky", 15)
	fmt.Println(person.name, person.age) // Miky 15
}

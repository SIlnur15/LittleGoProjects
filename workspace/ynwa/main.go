package main

type Animal interface {
	MakeSound() string
}

type Dog struct{}

func (d Dog) MakeSound() string {
	return "Гав-гав"
}

type Cat struct{}

func (c Cat) MakeSound() string {
	return "Мяу"
}

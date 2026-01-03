package main

import "fmt"

func changePointer(pp **int, newValue int) {
	*pp = &newValue      // разыменовываем pp и присваиваем ему адрес переменной newValue
	fmt.Println(pp, *pp) // 0xc000070030 0xc000010100 - оба вывода - адреса
}

func main() {
	x := 10
	ptr := &x // ptr — указатель на x (т.е. хранит адрес переменной x)

	y := 20
	changePointer(&ptr, y) // в функцию передаём: адрес указателя ptr и значение 20

	fmt.Println(*ptr) // 20
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	// Функция, которая будет выполнена только один раз.
	greet := func() {
		fmt.Println("Hello, World!")
	}

	// Вызов метода Do для выполнения функции greet.
	once.Do(greet)

	// Повторный вызов метода Do не приведет к повторному выполнению функции greet.
	once.Do(greet)
}

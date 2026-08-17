package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// Добавление пары ключ-значение
	m.Store("greeting", "hello")

	// Получение значения по ключу
	if value, ok := m.Load("greeting"); ok {
		fmt.Println(value.(string)) // -> hello
	}

	// Получение значения по отсутствующему ключу
	if _, ok := m.Load("nonexistent"); !ok {
		fmt.Println("значение не существует") // -> значение не существует
	}

	// Замена значения по ключу
	m.Store("greeting", "hi")

	// Удаление пары ключ-значение по ключу
	m.Delete("greeting")

	// Перебор всех пар ключ-значение и вывод их на экран
	m.Store("key1", "value1")
	m.Store("key2", "value2")
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key.(string) + ": " + value.(string))
		return true // продолжаем итерацию
	})

	// Замена значения по ключу (аналогично Store, но возвращает предыдущее значение, если оно существует)
	previous, loaded := m.Swap("key3", "value3")
	if loaded {
		fmt.Println(previous.(string)) // -> value1 (предыдущее значение ключа key3, если оно существует)
	}
}

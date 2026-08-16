package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var score int32 = 100 // Начальное значение

	success := atomic.CompareAndSwapInt32(&score, 100, 150) // Успешный случай: Проверяем: "Если score равен 100, то сделай его 150"
	fmt.Println("Успешно?", success)                        // Выведет: true
	fmt.Println("Score:", score)                            // Выведет: 150

	success2 := atomic.CompareAndSwapInt32(&score, 100, 200) // Неуспешный случай: Проверяем: "Если score равен 100, то сделай его 200". Но score уже равен 150! Условие не выполнено.
	fmt.Println("Успешно?", success2)                        // Выведет: false
	fmt.Println("Score:", score)                             // Выведет: 150 (значение НЕ изменилось)
}

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	switch num {
	case 2:
		fmt.Println("Неудовлетворительно")
	case 3:
		fmt.Println("Удовлетворительно")
	case 4:
		fmt.Println("Хорошо")
	case 5:
		fmt.Println("Отлично")
	default:
		fmt.Println("Некорректное значение!")
	}
}

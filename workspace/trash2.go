package main

import (
	"fmt"
	"math"
)

const (
	Even     = "even"
	Odd      = "odd"
	Positive = "positive"
	Negative = "negative"
	Squares  = "squares"
	Prime    = "prime"
	Factor10 = "factor10"
)

func filter(slice []int, filterFunc func(int) bool) []int {
	// менять эту функцию не нужно!!!
	// но я советую вам внимательно изучить эту функцию для саморазвития!

	// эта функция проходится по каждому элементу слайса
	// и отфильтровывает те, для которых filterFunc возвращает false

	result := make([]int, 0, len(slice))

	for _, num := range slice {
		if filterFunc(num) {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	// дополните эту мапу своими анонимными функциями!
	filters := map[string]func(int) bool{
		Even: func(numb int) bool { // Только чётные числа
			return numb%2 == 0
		},
		Odd: func(numb int) bool { // Только нечётные числа
			return numb%2 != 0
		},
		Positive: func(numb int) bool { // Только положительные числа
			return numb > 0
		},
		Negative: func(numb int) bool { // Только отрицательные числа
			return numb < 0
		},
		Squares: func(numb int) bool { // Только числа, которые являются квадратами
			if numb < 0 {
				return false
			}
			sqrt := math.Sqrt(float64(numb))
			return sqrt == float64(int(sqrt))
		},
		Prime: func(numb int) bool { // Только простые числа
			if numb <= 1 {
				return false
			}
			for i := 2; i*i <= numb; i++ {
				if numb%i == 0 {
					return false
				}
			}
			return true
		},
		Factor10: func(numb int) bool { // Только числа, которые кратны 10
			return numb%10 == 0
		},
	}

	// Ниже ничего менять не нужно
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	var operation string
	fmt.Scan(&operation)

	if filterFunc, exists := filters[operation]; exists {
		result := filter(numbers, filterFunc)
		for _, num := range result {
			fmt.Print(num, " ")
		}
	}
}

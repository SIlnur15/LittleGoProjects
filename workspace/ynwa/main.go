package main

import (
	"fmt"
	"math"
)

func safeAdd32(a, b int32) (int32, bool) {
	if b > 0 && a > math.MaxInt32-b {
		return 0, false // Переполнение
	}
	if b < 0 && a < math.MinInt32-b {
		return 0, false // Потеря значимости
	}
	return a + b, true
}

func main() {
	fmt.Println(safeAdd32(5, 15))
}

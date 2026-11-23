package benchmarking

import (
	"math/rand"
	"time"
)

// Функция с высокой CPU-нагрузкой
func heavyCalculation() int {
	result := 0
	for i := 0; i < 1000000; i++ {
		result += rand.Intn(100)
	}
	return result
}

func main() {
	for {
		heavyCalculation()
		time.Sleep(100 * time.Millisecond)
	}
}

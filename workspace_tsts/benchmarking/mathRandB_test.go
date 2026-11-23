package benchmarking

import (
	"testing"
)

// Функция с высокой CPU-нагрузкой
func BenchmarkHeavyCalculation(b *testing.B) {
	for b.Loop() {
		heavyCalculation()
	}
}


package main

import "testing"

func Add(x, y int) int {
	return x + y
}

func BenchmarkAdd(b *testing.B) {
	// Современный подход в Go 1.24+
	for b.Loop() {
		Add(1, 2)
	}
}

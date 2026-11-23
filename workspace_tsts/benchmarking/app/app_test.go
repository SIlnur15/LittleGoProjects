package app

import (
	"testing"
)

func BenchmarkAllocateMemory(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		allocateMemory()
	}
}

func BenchmarkUseCPU(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		useCPU()
	}
}

func BenchmarkUseChannels(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		useChannels()
	}
}

func BenchmarkPerformBusinessLogic(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		performBusinessLogic()
	}
}

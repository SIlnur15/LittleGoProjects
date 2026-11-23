package app

import (
	"log"
	"net/http"
	_ "net/http/pprof" // Автоматически регистрирует handlers на /debug/pprof
	"time"
)

func startProfilingServer() {
	// Запускаем HTTP-сервер для профилирования
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

func main() {
	startProfilingServer()

	// Ваше основное приложение
	for {
		performBusinessLogic()
		time.Sleep(time.Second)
	}
}

func performBusinessLogic() {
	// Имитация бизнес-логики
	allocateMemory()
	useCPU()
	useChannels()
}

func allocateMemory() {
	// Аллокация памяти
	data := make([]byte, 1024*1024) // 1MB
	for i := range data {
		data[i] = byte(i % 256)
	}
	time.Sleep(10 * time.Millisecond)
}

func useCPU() {
	// Использование CPU
	result := 0
	for i := 0; i < 1000000; i++ {
		result += i * i
	}
}

func useChannels() {
	// Использование каналов (потенциальные блокировки)
	ch := make(chan int, 1)
	ch <- 42
	<-ch
}

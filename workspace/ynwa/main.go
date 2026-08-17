package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	cond    *sync.Cond = sync.NewCond(&mu)
	counter int
)

func main() {
	for i := 0; i < 3; i++ {
		go listener(i)
	}

	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		mu.Lock()
		counter++
		fmt.Println("Counter:", counter)
		cond.Signal() // отправляем сигнал одной горутине
		mu.Unlock()
	}
}

func listener(id int) {
	for {
		mu.Lock()
		cond.Wait()
		fmt.Println("Listener:", id, " Counter:", counter)
		mu.Unlock()
	}
}

package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	counter int
)

func increment() {
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

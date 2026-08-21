package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var n int
	fmt.Scan(&n)
	var balance int64 = 0
	var wg sync.WaitGroup

	wg.Add(n)

	for i := 0; i < n; i++ {
		go add(&balance, &wg)
	}

	wg.Wait()
	fmt.Println("Баланс:", balance)
}

func add(balance *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(balance, 3)
}

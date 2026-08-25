package main

import (
	"fmt"
	"sync"
)

func main() {
	const n = 5
	results := make([]string, n)
	var wg sync.WaitGroup

	for i := range n {
		wg.Go(func() {
			results[i] = fmt.Sprintf("Привет из горутины %d", i) // пишем в свой индекс
		})
	}

	wg.Wait()

	for _, msg := range results {
		fmt.Println(msg)
	}
	fmt.Println("Все горутины завершились!")
}

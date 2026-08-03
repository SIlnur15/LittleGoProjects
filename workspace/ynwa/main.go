package main

import (
	"fmt"
	"sync"
)

func main() {
	oneWriterManyReaders()
}

func oneWriterManyReaders() {
	var wg sync.WaitGroup
	ch := make(chan int)
	const readers = 3

	wg.Add(1)
	go func() { // Единственный писатель
		defer wg.Done()
		for i := 0; i < readers; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i < readers; i++ { // Множество читателей
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for val := range ch {
				fmt.Printf("Reader %d got %d\n", id, val)
			}
		}(i)
	}
	wg.Wait()
}

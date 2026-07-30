package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%dth goroutineis printed\n", i)
		}()
	}
	wg.Wait()
	fmt.Println("all goroutines are printed")
}

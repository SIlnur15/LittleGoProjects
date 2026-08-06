package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	input := make(chan string)
	output1 := make(chan string)
	output2 := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for s := range input {
			words := strings.Split(s, " ")
			for i, w := range words {
				words[i] = strings.ToUpper(w)
			}
			output1 <- strings.Join(words, " ")
		}
		close(output1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for s := range output1 {
			vowels := "AEIOUaeiou"
			for _, v := range vowels {
				s = strings.Replace(s, string(v), "*", -1)
			}
			output2 <- s
		}
		close(output2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for s := range output2 {
			fmt.Println(s)
		}
	}()

	input <- "Hello world"
	input <- "Go is awesome"
	input <- "Concurrency is the future"
	close(input)

	wg.Wait()
}

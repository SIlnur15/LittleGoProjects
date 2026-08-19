package main

import (
	"fmt"
	"sync"
)

// joinChannels принимает несколько read-only каналов и объединяет их в один.
func joinChannels(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				mergedCh <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}

// sendAll отправляет все значения из nums в ch, затем закрывает его.
func sendAll(ch chan<- int, nums []int) {
	defer close(ch)
	for _, n := range nums {
		ch <- n
	}
}

func main() {
	// Данные вынесены: горутины только отправляют, не хранят
	dataSets := [][]int{
		{1, 2, 3},
		{20, 10, 30},
		{300, 200, 100},
	}

	channels := make([]chan int, len(dataSets))
	for i := range channels {
		channels[i] = make(chan int)
	}

	for i, ch := range channels {
		go sendAll(ch, dataSets[i])
	}

	// Конвертируем []chan int → []<-chan int для joinChannels
	readOnly := make([]<-chan int, len(channels))
	for i, ch := range channels {
		readOnly[i] = ch
	}

	for num := range joinChannels(readOnly...) {
		fmt.Println(num)
	}
}

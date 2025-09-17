package main

import (
	"fmt"
	"math"
)

func main() {
	var evens, not_evens []int
	var val int

	for i := 0; i < 15; i++ {
		fmt.Scan(&val)
		if val%2 == 0 {
			evens = append(evens, val)
		} else {
			not_evens = append(not_evens, val)
		}
	}
	length := int(math.Max(float64(len(evens)), float64(len(not_evens))))

	for i := 0; i < length; i++ {
		if i < len(evens) {
			fmt.Printf("%d ", evens[i])
		}
		if i < len(not_evens) {
			fmt.Printf("%d ", not_evens[i])
		}
	}
}

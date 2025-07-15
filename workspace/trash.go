package main

import "fmt"

func main() {
	var counter float32
	massive1 := [5] float32 {1.2, 2.2, 2.1, 3.2, 3.1}
	for _, value := range massive1 {
		counter += value
	}
	output := counter/float32(len(massive1))
	fmt.Printf("%.1f\n", output)
}


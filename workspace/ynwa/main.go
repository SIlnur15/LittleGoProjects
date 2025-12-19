package main

import "fmt"

func main() {
	var (
		arr [10]int
		sum int
	)
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < len(arr); i = i + 2 {
		sum += arr[i]
	}
	fmt.Println(sum)
}

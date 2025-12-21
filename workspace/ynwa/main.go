package main

import "fmt"

func main() {
	var (
		n   int
		inp int
		arr []int
	)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&inp)
		if inp%2 == 0 {
			arr = append(arr, inp)
		}
	}
	fmt.Println(arr)
}

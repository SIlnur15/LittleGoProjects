package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scan(&n)
	slc := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slc[i])
	}
	slices.Reverse(slc)
	for idx := range slc {
		fmt.Print(slc[idx], " ")
	}
}

package main

import (
	"fmt"
	"slices"
)

func union(a, b []int) []int {
	for i := 0; i < len(b); i++ {
		a = append(a, b[i])
	}
	slices.Sort(a)
	slices.Compact(a)
	b = make([]int, 0)
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			b = append(b, a[i])
		}
	}
	return b
}

func main() {
	var (
		a, b, value int
	)
	fmt.Scan(&a, &b)
	arr1 := make([]int, a)
	arr2 := make([]int, b)
	for i := 0; i < a; i++ {
		fmt.Scan(&value)
		arr1[i] = value
	}
	for i := 0; i < b; i++ {
		fmt.Scan(&value)
		arr2[i] = value
	}
	fmt.Println(union(arr1, arr2))
}

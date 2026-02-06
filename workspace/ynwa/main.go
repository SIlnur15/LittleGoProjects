package main

import (
	"errors"
	"fmt"
)

func getValue(arr []int, idx int) (int, error) {
	if idx >= len(arr) || idx < 0 {
		return 0, errors.New("list index out of range")
	} else {
		return arr[idx], nil
	}
}

func main() {
	var (
		idx int
	)
	arr := []int{5, 15, 25}
	fmt.Scan(&idx)
	fmt.Println(getValue(arr, idx))

}

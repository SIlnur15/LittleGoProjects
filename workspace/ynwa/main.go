package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		n          int
		arr1, arr2 []int
	)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := scanner.Text()
		input_int, _ := strconv.Atoi(input)
		arr1[i] = input_int
	}
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := scanner.Text()
		input_int, _ := strconv.Atoi(input)
		arr2[i] = input_int
	}
	fmt.Println(arr1, arr2)
}

package main

import "fmt"

func main() {
	arr := [7]int{4, 2, 9, 5, 6, 7}
	var first, second int
	fmt.Scan(&first, &second)
	fmt.Println(arr[first], arr[second])
}

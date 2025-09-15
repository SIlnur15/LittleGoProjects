package main

import "fmt"

func main() {
	res := 0
	nums := [5]int{1, 2, 3, 4, 5}
	for i, v := range nums {
		if i%2 != 0 {
			res += v
		}
	}
	fmt.Println(res)
}

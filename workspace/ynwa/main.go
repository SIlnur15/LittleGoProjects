package main

import "fmt"

func main() {
	var num, sum int
	for {
		fmt.Scan(&num)
		sum += num
		if num == 0 {
			break
		}
	}
	fmt.Println(sum)
}

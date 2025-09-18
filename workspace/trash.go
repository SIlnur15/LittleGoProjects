package main

import "fmt"

func main() {
	var values, value int
	mp := make(map[int]int)
	fmt.Scan(&values)
	for i := 0; i < values; i++ {
		fmt.Scan(&value)
		mp[value] += 1
	}
	fmt.Println(mp)
}

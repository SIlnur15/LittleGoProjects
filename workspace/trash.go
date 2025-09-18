package main

import "fmt"

func main() {
	slice := []int{5, 5, 5, 5, 5}
	fmt.Println(cap(slice)) // 5
	slice2 := []int{3, 3, 3}
	fmt.Println(cap(slice2)) // 3
}

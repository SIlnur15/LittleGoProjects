package main

import "fmt"

func Identity[T any](v T) T {
	return v
}

func main() {
	fmt.Println(Identity(42))      // int
	fmt.Println(Identity("hello")) // string
}

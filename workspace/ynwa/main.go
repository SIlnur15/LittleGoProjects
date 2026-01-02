package main

import "fmt"

func main() {
	a := 100
	ptr := &a
	ptr2 := new(int)
	ptr2 = ptr
	*ptr = 200
	fmt.Println(*ptr, *ptr2)
}

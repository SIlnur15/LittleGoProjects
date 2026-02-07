package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i interface{}
	fmt.Println(unsafe.Sizeof(i))
	var j struct{}
	fmt.Println(unsafe.Sizeof(j))
}

package main

import "fmt"

var num int16 = 15

func Double(numPointer *int16) {
	*numPointer = *numPointer*2
	fmt.Println(*numPointer)
}

func main() {
	Double(&num)
	fmt.Println(num)
}

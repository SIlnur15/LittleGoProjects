package main

import ("fmt"; "reflect")

func main() {
	var myInt int
	var myIntPointer *int
	myInt = 42
	myIntPointer = &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer) 						// 42
	fmt.Println(reflect.TypeOf(myIntPointer)) 		// *int
	fmt.Println(reflect.TypeOf(*myIntPointer)) 		// int
}
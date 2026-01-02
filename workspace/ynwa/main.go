package main

import "fmt"

func dubleMaker(numb *int) {
	*numb = *numb * 2
}

func main() {
	a := 10
	fmt.Println(a)
	dubleMaker(&a)
	fmt.Println(a)
}

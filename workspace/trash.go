package main

import "fmt"

type myType struct {
	name   string
	counts int
	flag   bool
}

func main() {
	var special1 myType
	special1.counts = 25
	special1.name = "number512"
	fmt.Println(special1)
	fmt.Println(special1.counts)
}

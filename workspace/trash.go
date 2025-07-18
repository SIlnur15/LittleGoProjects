package main

import "fmt"

type myType struct {
	name  string
	value int
}

func expreience(param *int) {
	*param += 2
}

func main() {
	var man myType
	fmt.Println(man.value)
	expreience(&man.value)
	fmt.Println(man.value)
}

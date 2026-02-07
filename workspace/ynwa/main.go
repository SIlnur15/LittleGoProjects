package main

import "fmt"

type myPerson struct {
	name string
	age  int
}

func main() {
	h := &myPerson{"Haid", 95}
	s := myPerson{"Said", 87}
	addrS := &s
	fmt.Println(h, s)
	fmt.Println(&h, &addrS)
	h.name = "Haidaric"
	s.name = "Saidaric"
	fmt.Println(h.name, s.name)
	fmt.Println(&h, *addrS)
}

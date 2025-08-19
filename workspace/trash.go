package main

import "fmt"

func main() {
	i := 1
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		// fallthrough
	case 3:
		fmt.Println("Bye-bye")
	}
}

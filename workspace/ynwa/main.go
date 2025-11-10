package main

import "fmt"

func main() {
	num := 123.456789
	fmt.Printf("%%f:     '%f'\n", num)    // стандарт
	fmt.Printf("%%9f:    '%9f'\n", num)   // ширина 9
	fmt.Printf("%%.2f:   '%.2f'\n", num)  // точность 2
	fmt.Printf("%%9.2f:  '%9.2f'\n", num) // ширина 9, точность 2
	fmt.Printf("%%9.f:   '%9.f'\n", num)  // ширина 9, точность 0
	fmt.Printf("%%09.f:  '%09.f'\n", num) // нули, ширина 9, точность 0
}

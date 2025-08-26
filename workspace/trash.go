package main

import "fmt"

func main() {
	var code string
	fmt.Scan(&code)
	switch code {
	case "en":
		fmt.Println("English")
	case "fr":
		fmt.Println("French")
	case "ru", "rus":
		fmt.Println("Russian")
	default:
		fmt.Println("Unknown")
	}
}

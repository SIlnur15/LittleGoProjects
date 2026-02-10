package main

import "fmt"

type Callback func(string)

func notify(cb Callback, message string) {
	cb(message)
}

func main() {
	notify(func(msg string) {
		fmt.Println("statement:", msg)
	}, "hello")
}

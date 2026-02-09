package main

import "fmt"

type account struct {
	balance float32
}

func main() {
	accounts := []account{
		{balance: 100},
		{balance: 200},
		{balance: 300},
	}
	for _, a := range accounts {
		a.balance += 1000
	}
	for idx, _ := range accounts {
		accounts[idx].balance += 1000
	}
	fmt.Println(accounts)
}

package main

import (
	"dates"
	"fmt"
	"log"
)

func main() {
	interest := dates.Dates{}
	err := interest.SetDay(22)
	if err != nil {
		log.Fatal(err)
	}
	err = interest.SetMonth(2)
	if err != nil {
		log.Fatal(err)
	}
	err = interest.SetYear(1994)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(interest)
	fmt.Println(interest.Day())
	fmt.Println(interest.Month())
	fmt.Println(interest.Year())
}

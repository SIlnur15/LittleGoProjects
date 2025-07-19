package main

import (
	"calendar"
	"fmt"
)

func main() {
	var value calendar.Date
	value.SetDay(22)
	value.SetMonth(2)
	value.SetYear(1994)
	fmt.Println(value)
}

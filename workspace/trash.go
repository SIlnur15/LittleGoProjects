package main

import (
	"fmt"
)

func main() {
	var rows_count, column_count int
	fmt.Println("please enter rows' amount")
	fmt.Scan(&rows_count)
	fmt.Println("please enter columns' number")
	fmt.Scan(&column_count)
	table := make([][]int, rows_count) // задаю двумерный срез
	// the first conditions//
	for i := 0; i < len(table); i++ {
		fmt.Printf("%v, type %T\n", table[i], table[i])
	}
	// end of block with the first conditions
	for i := 0; i < len(table); i++ {
		table[i] = make([]int, column_count) // можно задавать только одномерный срез
		for j := 0; j < column_count; j++ {
			fmt.Println("please enter value in column")
			fmt.Scan(&table[i][j])
		}
	}
	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
	fmt.Println(table[2][1])
}

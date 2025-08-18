package main

import (
	"fmt"
	"math"
)

func main() {
	var input1, input2, input3, input4 float64             // ф-я math.Pow требует такой тип
	_, err := fmt.Scan(&input1, &input2, &input3, &input4) // считываем числа из os.Stdin
	if err != nil {
		fmt.Println("wrong entering:", err)
		return
	}
	under_sqrt := math.Pow(input1-input3, 2) + math.Pow(input2-input4, 2)
	d := math.Pow(under_sqrt, 0.5) // attention 1/2 не покатит !!!
	fmt.Println(d)
}

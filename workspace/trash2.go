package main

import "fmt"

func main() {
	// место для ваших анонимных функций
	calculator := map[string]func(int, int) int{
		"+": func(c, d int) int {
			return c + d
		},

		"-": func(c, d int) int {
			return c - d
		},

		"*": func(c, d int) int {
			return c * d
		},

		"/": func(c, d int) int {
			if d != 0 {
				return c / d
			} else {
                fmt.Println("Делить на ноль нельзя!")
                return 0
            }
		},
	}

	// ниже можете ничего не менять
	var operation string
	fmt.Scan(&operation)

	var a, b int
	fmt.Scan(&a, &b)

	result := calculator[operation](a, b)

	fmt.Println(result)
}

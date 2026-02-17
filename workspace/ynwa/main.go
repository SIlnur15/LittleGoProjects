package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("test.txt")
	reader := bufio.NewReader(file)

	bytes, _ := reader.Peek(5)

	fmt.Printf("%s\n", bytes) // выведет первые 5 байт файла
}

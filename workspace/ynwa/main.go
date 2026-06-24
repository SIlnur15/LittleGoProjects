package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args содержит все аргументы командной строки,
	// включая имя исполняемого файла
	args := os.Args
	fmt.Println(args)
}

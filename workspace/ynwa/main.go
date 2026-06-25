package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Определение флагов
	verbose := flag.Bool("v", false, "verbose output")
	port := flag.Int("port", 8080, "server port")

	// Парсинг флагов
	flag.Parse()

	// Аргументы командной строки (не флаги)
	args := flag.Args()

	fmt.Printf("Verbose: %v\n", *verbose)
	fmt.Printf("Port: %d\n", *port)
	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("All command line args (including flags): %v\n", os.Args)
}

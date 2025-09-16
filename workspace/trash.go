package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	output := input.Text()
	fmt.Println(output)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("please enter your name: ")
	littleScanner := bufio.NewScanner(os.Stdin)
	littleScanner.Scan()
	name := littleScanner.Text()
	fmt.Printf("So, nice to meet you, %s!", name)
}

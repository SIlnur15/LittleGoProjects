package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	new_scanner := bufio.NewScanner(os.Stdin)
	for new_scanner.Scan() {
		name := new_scanner.Text()
		fmt.Print(name)
	}
}

// func main() {
// 	new_scanner := bufio.NewScanner(os.Stdin)
// 	for new_scanner.Scan() {
// 		name := new_scanner.Text()
// 		fmt.Print(name)
// 	}
// }

// func main() {
// 	new_scanner := bufio.NewScanner(os.Stdin)
// 	for new_scanner.Scan() {
// 		name := new_scanner.Text()
// 		fmt.Print(name)
// 	}
// }

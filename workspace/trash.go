package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	delimiter, _ := reader.ReadString('\n')
	delimiter = strings.TrimSpace(delimiter)
	s1, _ := reader.ReadString('\n')
	s1 = strings.TrimSpace(s1)
	s2, _ := reader.ReadString('\n')
	s2 = strings.TrimSpace(s2)
	s3, _ := reader.ReadString('\n')
	s3 = strings.TrimSpace(s3)
	fmt.Print(s1, delimiter, s2, delimiter, s3)
}

// func main() {
// 	new_scanner := bufio.NewScanner(os.Stdin)
// 	for new_scanner.Scan() {
// 		name := new_scanner.Text()
// 		fmt.Print(name)
// 	}
// }

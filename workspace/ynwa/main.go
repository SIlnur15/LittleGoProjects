package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content, err := os.ReadFile("stepik_io_1.text")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	input := string(content)
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(input, -1)
	sum := 0
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		sum += num
	}
	fmt.Println(sum)
}

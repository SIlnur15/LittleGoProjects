package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(arr ...int) int {
	mx := 0
	for _, elem := range arr {
		if elem > mx {
			mx = elem
		}
	}
	return mx
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	kaput := strings.Fields(input)
	kaputInt := make([]int, len(kaput))
	for idx, _ := range kaput {
		kaputIntElem, _ := strconv.Atoi(kaput[idx])
		kaputInt[idx] = kaputIntElem
	}
	fmt.Println(max(kaputInt...))
}

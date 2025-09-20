package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mp := make(map[int]bool, 1)
	result := make([]int, 0)
	newReader := bufio.NewReader(os.Stdin)
	input1, _ := newReader.ReadString('\n')
	input1 = strings.Trim(input1, "\n")
	mass1 := strings.Split(input1, " ")
	for idx, _ := range mass1 {
		value, _ := strconv.Atoi(mass1[idx])
		mp[value] = true
	}
	input2, _ := newReader.ReadString('\n')
	input2 = strings.Trim(input2, "\n")
	mass2 := strings.Split(input2, " ")
	for idx, _ := range mass2 {
		value, _ := strconv.Atoi(mass2[idx])
		_, ok := mp[value]
		if ok {
			result = append(result, value)
		}
	}
	for _, value := range result {
		fmt.Printf("%d ", value)
	}
}

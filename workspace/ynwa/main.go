package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var output map[string]interface{}

	file, err := os.Open("stepik4-4.json")
	if err != nil {
		fmt.Println("trouble")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&output)
	if _, ok := output["people"].([]interface{}); ok {
		fmt.Println("kuku")
	}
}

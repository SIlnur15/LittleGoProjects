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
	if people, ok := output["people"].([]interface{}); ok {
		for _, person := range people {
			if gender, ok := person.(map[string]interface{}); ok {
				if gender["gender"] == "Female" {
					fmt.Printf("%s %s\n", gender["first_name"], gender["last_name"])
				}
			}
		}
	}
}

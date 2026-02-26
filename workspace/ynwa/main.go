package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var (
		output map[string]interface{}
		summa  float64
		lenArr float64
	)

	file, err := os.Open("stepik4-4-2.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&output)
	if err != nil {
		fmt.Println("2. trouble")
	}
	if students, ok := output["students"].([]interface{}); ok {
		for _, student := range students {
			if studentStudent, ok := student.(map[string]interface{}); ok {
				if person, ok := studentStudent["Rating"].(interface{}); ok {
					if arr, ok := person.([]interface{}); ok {
						for _, elem := range arr {
							if elemElem, ok := elem.(float64); ok {
								summa += elemElem
							}
							lenArr += 1
						}
					}
				}
			}
		}
		fmt.Printf("%.2f\n", summa/lenArr)
	}
}

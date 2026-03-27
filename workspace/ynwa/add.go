package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	sum := 0
	req, err := http.Get("https://open.er-api.com/v6/latest/USD")
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()

	var rts map[string]interface{}

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&rts)
	if err != nil {
		fmt.Println(err)
	}
	if Rates, ok := rts["rates"].(map[string]interface{}); ok {
		for valute := range Rates {
			if string(valute[0]) == "T" {
				sum += 1
			}
		}
	}
	fmt.Println(sum)
}

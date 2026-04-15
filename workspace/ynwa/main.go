package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://api.zippopotam.us/us/90210")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var cont map[string]interface{}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&cont)
	if err != nil {
		fmt.Println(err)
		return
	}

	if places, ok := cont["places"].([]interface{}); ok {
		if sample, ok := places[0].(map[string]interface{}); ok {
			fmt.Println(sample["place name"]) // Beverly Hills
		}
	}
}

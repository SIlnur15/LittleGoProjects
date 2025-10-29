package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var myVal map[string]any
	bte := `{"some":"json"}`
	err := json.Unmarshal([]byte(bte), &myVal)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myVal) // map[some:json]
}

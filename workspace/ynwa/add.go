package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := []byte(`{"name":"John","age":30,"city":"New York"`)
	isValid := json.Valid(jsonData)
	fmt.Println(isValid)
}

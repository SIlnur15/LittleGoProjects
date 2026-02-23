package main

import (
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	Name string `json:"first_name"`
}

func main() {
	jsonData := []byte(`{"Name":"John","first_name":"Fred"}`)
	var myStruct MyStruct
	err := json.Unmarshal([]byte(jsonData), &myStruct)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(myStruct.Name)
}

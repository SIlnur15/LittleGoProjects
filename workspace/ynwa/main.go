package main

import (
	"encoding/json"
	"fmt"
)

type MyPerson struct {
	Name string `json:"myName`
	Age  int    `json:"myAge"`
}

func main() {
	p := MyPerson{"Haidaric", 75}
	jsonProd, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", jsonProd)
}

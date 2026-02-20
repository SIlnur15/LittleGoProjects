package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"sweetName"`
	Age  int    `json:"littleAge"`
}

func main() {
	p := Person{Name: "Haidaric", Age: 75}
	per, err := json.Marshal(p)
	if err != nil {
		fmt.Println("there is a trouble")
		return
	}
	fmt.Printf("%s\n", per) // {"name":"Haidaric","age":75}
}

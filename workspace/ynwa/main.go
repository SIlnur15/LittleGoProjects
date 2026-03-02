package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	// создаем объект запроса со значением body
	requestBody := []byte(`{"foo": "bar"}`)

	req, err := http.NewRequest(
		"POST",
		"https://jsonplaceholder.typicode.com/posts",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}

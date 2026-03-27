package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	cli := &http.Client{
		Timeout: time.Second * 5, // Wait no more than 5 seconds for an answer
	}
	req, err := http.NewRequest(
		"GET",
		"https://api.chucknorris.io/jokes/rBOPF3bYRN6S6P2Y2ZCCWw",
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("User-Agent", "Golang-Client") // to bypass (обход) bot protection

	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}

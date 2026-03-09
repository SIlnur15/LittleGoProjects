package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the about page.")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/about", aboutHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}
}

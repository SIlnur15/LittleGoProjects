package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	http.Handle("/", &myHandler{})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}
}

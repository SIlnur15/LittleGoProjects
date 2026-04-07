package main

import (
	"fmt"
	"net/http"
)

func helloEnHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func helloRuHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет, мир!"))
}

func helloEsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola, Mundo!"))
}

func helloJpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("「こんにちは世界」"))
}

func main() {
	http.HandleFunc("/en", helloEnHandler)
	http.HandleFunc("/ru", helloRuHandler)
	http.HandleFunc("/es", helloEsHandler)
	http.HandleFunc("/jp", helloJpHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}
}

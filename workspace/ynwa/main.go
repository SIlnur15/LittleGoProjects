package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Получаем все query parameters как map
	queryParams := r.URL.Query()

	// Получаем конкретный параметр
	name := queryParams.Get("name")
	age := queryParams.Get("age")

	// Если параметр может встречаться несколько раз
	ids := queryParams["id"] // возвращает []string

	fmt.Fprintf(w, "Name: %s, Age: %s, IDs: %v", name, age, ids)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

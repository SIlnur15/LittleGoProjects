package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем все query-параметры
	queryParams := r.URL.Query()

	// Выводим всю карту параметров
	fmt.Fprintf(w, "Все параметры: %v\n\n", queryParams)

	// Получаем отдельные значения
	q := queryParams.Get("q")
	page := queryParams.Get("page")
	sort := queryParams.Get("sort")

	fmt.Fprintf(w, "q = %s\n", q)
	fmt.Fprintf(w, "page = %s\n", page)
	fmt.Fprintf(w, "sort = %s\n", sort)

	// Обрабатываем параметр с несколькими значениями
	tags := queryParams["tags"]
	fmt.Fprintf(w, "tags = %v\n", tags)
}

func main() {
	http.HandleFunc("/search", handler)
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

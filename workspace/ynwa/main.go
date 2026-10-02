package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("APP_ENV")
	fmt.Fprintf(w, "Приложение работает! Среда: %s", env)
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	fmt.Println("Сервер запущен на порту", port)
	http.ListenAndServe(":"+port, nil)
}

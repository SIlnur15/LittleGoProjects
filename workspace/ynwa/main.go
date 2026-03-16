package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Создаем свой собственный экземпляр роутера (мультиплексора)
	mux := http.NewServeMux()

	// 2. Регистрируем маршруты именно в нашем локальном mux
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Главная страница")
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "О нас")
	})

	// 3. Передаем наш mux вторым аргументом.
	// Теперь ListenAndServe не использует глобальный DefaultServeMux.
	fmt.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", mux)
}

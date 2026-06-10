package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title   string
	Content string
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Главная страница",
		Content: "Добро пожаловать на наш сайт!",
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

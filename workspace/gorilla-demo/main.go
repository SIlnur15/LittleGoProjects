package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Title  string `json:"title,omitempty"`
	Rating int    `json:"rating,omitempty"`
}

var movies = []Movie{
	{"Побег", 9},
	{"The Godfather", 9},
} // Создадим слайс с тестовыми данными

func GetMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Title == params["title"] {
			json.NewEncoder(w).Encode(&item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", GetMovies).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

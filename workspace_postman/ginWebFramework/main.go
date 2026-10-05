package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Movie представляет структуру фильма
type Movie struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
}

var movies = []Movie{
	{ID: 1, Title: "Самогонщики", Director: "Леонид Гайдай", Year: 1961},
	{ID: 2, Title: "Бриллиантовая рука", Director: "Леонид Гайдай", Year: 1968},
}

func main() {
	// Создаем новый экземпляр роутера
	r := gin.Default()

	// Получить список всех фильмов
	r.GET("/movies", func(c *gin.Context) {
		c.JSON(http.StatusOK, movies)
	})

	// Получить фильм по ID
	r.GET("/movies/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
			return
		}
		for _, m := range movies {
			if m.ID == id {
				c.JSON(http.StatusOK, m)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
	})

	// Добавить новый фильм
	r.POST("/movies", func(c *gin.Context) {
		var newMovie Movie
		if err := c.ShouldBindJSON(&newMovie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newMovie.ID = getNextID()
		movies = append(movies, newMovie)
		c.JSON(http.StatusCreated, newMovie)
	})

	// Обновить фильм по ID
	r.PUT("/movies/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
			return
		}
		var updatedMovie Movie
		if err := c.ShouldBindJSON(&updatedMovie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, m := range movies {
			if m.ID == id {
				updatedMovie.ID = id
				movies[i] = updatedMovie
				c.JSON(http.StatusOK, updatedMovie)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
	})

	// Удалить фильм по ID
	r.DELETE("/movies/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
			return
		}
		for i, m := range movies {
			if m.ID == id {
				movies = append(movies[:i], movies[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
	})

	r.Run(":8080") // Запуск сервера на порту 8080
}

func getNextID() int {
	maxID := 0
	for _, m := range movies {
		if m.ID > maxID {
			maxID = m.ID
		}
	}
	return maxID + 1
}

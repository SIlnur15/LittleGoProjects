package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Базовый пример middleware для логирования запросов
func LoggerMiddleware(c *gin.Context) {
	start := time.Now()
	c.Next() // Пропускаем запрос дальше
	// После обработки запроса
	latency := time.Since(start)
	log.Printf("[%s] %s | %v", c.Request.Method, c.Request.URL.Path, latency)
}

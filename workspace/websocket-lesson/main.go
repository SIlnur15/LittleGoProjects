package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

// Настраиваем Upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin позволяет контролировать, с каких доменов разрешены запросы.
	// Для обучения мы разрешим любые подключения.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// 1. Апгрейд соединения
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Ошибка при апгрейде: %v", err)
		return
	}
	// 2. Гарантируем закрытие соединения при выходе из функции
	defer conn.Close()

	log.Println("Клиент успешно подключен")

	// 3. Бесконечный цикл для чтения и записи сообщений
	for {
		// Читаем сообщение от клиента
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Ошибка при чтении: %v", err)
			break // Выходим из цикла, если клиент отключился или произошла ошибка
		}

		// Выводим полученное сообщение в консоль сервера
		log.Printf("Получено сообщение: %s", p)

		// Отправляем сообщение обратно (эхо)
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("Ошибка при записи: %v", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections) // Регистрируем маршрут
	log.Println("Сервер запущен на :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

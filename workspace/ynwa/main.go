package main

import (
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	// 1. Настраиваем менеджер сертификатов
	manager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example.com"),
		Cache:      autocert.DirCache("certs"),
	}

	// 2. Основной сервер (HTTPS) на порту 443
	server := &http.Server{
		Addr:      ":443",
		TLSConfig: manager.TLSConfig(), // Упрощенная настройка TLS
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Привет! Это защищенное соединение."))
		}),
	}

	// 3. Запускаем вспомогательный сервер (HTTP) на порту 80 для редиректа
	// manager.HTTPHandler(nil) автоматически перенаправляет на HTTPS
	go http.ListenAndServe(":80", manager.HTTPHandler(nil))

	// 4. Запускаем основной сервер
	// Либо используйте ListenAndServeTLS("", ""), либо просто ListenAndServe()
	// так как TLSConfig уже настроен выше
	server.ListenAndServeTLS("", "")
}

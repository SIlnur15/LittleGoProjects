package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// 1. Тестирование HTTP-обработчиков (Handlers)
func TestHandler(t *testing.T) {
	// Создаем запрос
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(YourHandlerFunction)

	// Вызываем обработчик
	handler.ServeHTTP(rr, req)

	// Проверяем статус код
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверяем тело ответа
	expected := `{"status":"ok"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// 2. Тестирование HTTP-клиентов
func TestClient(t *testing.T) {
    // Создаем тестовый сервер
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, `{"status":"ok"}`)
    }))
    defer ts.Close()

    // Делаем запрос к тестовому серверу
    res, err := http.Get(ts.URL)
    if err != nil {
        t.Fatal(err)
    }

    // Проверяем ответ
    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        t.Fatal(err)
    }

    expected := `{"status":"ok"}` + "\n"
    if string(body) != expected {
        t.Errorf("unexpected body: got %v want %v", string(body), expected)
    }
}

                  
// 3. TLS-сервер для тестирования HTTPS
ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, client")
}))
defer ts.Close()

client := ts.Client()
res, err := client.Get(ts.URL)
if err != nil {
    t.Fatal(err)
}
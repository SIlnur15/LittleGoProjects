package main

import (
	"context"
	"fmt"
	"log"
	"myproject/userpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// Подключаем сгенерированный из proto-файла код
)

func main() {
	conn, err := grpc.NewClient("192.168.1.50:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	// 1. Устанавливаем сетевое соединение с удаленным сервером

	// 2. Создаем заглушку (Stub).
	// Функция NewUserServiceClient — это сгенерированный фреймворком код.
	// Переменная client теперь содержит тот самый Stub.
	client := userpb.NewUserServiceClient(conn)

	// 3. Создаем контекст для контроля таймаута
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// ==========================================
	// 4. ВОТ ОН — ВЫЗОВ МЕТОДА-ЗАГЛУШКИ (Stub):
	// ==========================================
	resp, err := client.GetUser(ctx, &userpb.UserRequest{UserId: 42})
	// Для твоего кода это выглядит как обычный локальный метод структуры client.
	// Но внутри этого метода фреймворк прямо сейчас превратил ID 42 в байты,
	// отправил их по сети на IP 192.168.1.50, дождался ответа и вернул структуру resp!
	if err != nil {
		log.Fatalf("Ошибка при вызове удаленного метода: %v", err)
	}

	// Выводим результат так, будто функция выполнилась прямо здесь
	fmt.Printf("Пользователь найден! Имя: %s, Email: %s\n", resp.Name, resp.Email)
}

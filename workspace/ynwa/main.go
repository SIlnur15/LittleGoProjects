package main

import (
	"context"
	"fmt"
	"log"
	"myproject/userpb"
	"net"

	"google.golang.org/grpc"
)

// 1. Создаем структуру нашего сервера.
// Она ОБЯЗАТЕЛЬНО должна включать в себя UnimplementedUserServiceServer.
// Это нужно для обратной совместимости, если в будущем контракт (.proto) изменится.
type MyUserServer struct {
	userpb.UnimplementedUserServiceServer
}

func (cc *ClientConn) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...CallOption) error

// 2. Реализуем метод GetUser, который мы описали в .proto файле.
// Именно эту функцию клиент вызывает через свою заглушку (Stub).
func (s *MyUserServer) GetUser(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	// Логируем входящий запрос. req.UserId уже автоматически распакован из байтов!
	fmt.Printf("Сервер: Получен запрос для User ID: %d\n", req.UserId)

	// Имитируем поиск в базе данных.
	// В реальном проекте здесь был бы SQL-запрос вроде: SELECT name, email FROM users WHERE id = ...
	if req.UserId == 42 {
		return &userpb.UserResponse{
			Name:  "Алексей Гофер",
			Email: "alex@golang.org",
		}, nil
	}

	// Если пользователь не найден, возвращаем пустой ответ или ошибку gRPC
	return &userpb.UserResponse{Name: "Неизвестный пользователь"}, nil
}

func main() {
	// 3. Открываем TCP-порт для прослушивания входящих соединений
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Не удалось открыть порт: %v", err)
	}

	// 4. Создаем экземпляр gRPC-сервера
	grpcServer := grpc.NewServer()

	// 5. Регистрируем наш сервер MyUserServer внутри gRPC-фреймворка
	userpb.RegisterUserServiceServer(grpcServer, &MyUserServer{})

	fmt.Println("gRPC сервер успешно запущен на порту :50051...")

	// 6. Запускаем бесконечный цикл обработки запросов
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка при работе сервера: %v", err)
	}
}

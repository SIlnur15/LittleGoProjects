package main

import (
	"fmt"
	"time"
)

// Order представляет структуру заказа
type Order struct {
	ID     int
	Amount float64
	Valid  bool
}

func main() {
	// Создаем каналы для каждого этапа
	receivedOrders := make(chan Order, 10)  // Принятые заказы
	validatedOrders := make(chan Order, 10) // Проверенные заказы
	processedOrders := make(chan Order, 10) // Обработанные заказы

	// Запускаем этапы обработки в отдельных горутинах
	go receiveOrders(receivedOrders)
	go validateOrders(receivedOrders, validatedOrders)
	go processOrders(validatedOrders, processedOrders)

	// Выводим результаты обработки
	for order := range processedOrders {
		fmt.Printf("Заказ %d завершен: сумма %.2f, валидность %v\n",
			order.ID, order.Amount, order.Valid)
	}
}

// receiveOrders - этап получения заказов
func receiveOrders(out chan<- Order) {
	for i := 1; i <= 5; i++ {
		order := Order{ID: i, Amount: float64(i) * 10}
		fmt.Printf("Получен заказ %d\n", order.ID)
		out <- order
		time.Sleep(time.Millisecond * 200) // Имитация работы
	}
	close(out)
}

// validateOrders - этап проверки заказов
func validateOrders(in <-chan Order, out chan<- Order) {
	for order := range in {
		// Простая проверка - сумма должна быть больше 20
		order.Valid = order.Amount > 20
		fmt.Printf("Проверка заказа %d: валидность %v\n", order.ID, order.Valid)
		out <- order
		time.Sleep(time.Millisecond * 300) // Имитация работы
	}
	close(out)
}

// processOrders - этап обработки заказов
func processOrders(in <-chan Order, out chan<- Order) {
	for order := range in {
		if order.Valid {
			fmt.Printf("Обработка заказа %d\n", order.ID)
		} else {
			fmt.Printf("Отмена заказа %d (невалидный)\n", order.ID)
		}
		out <- order
		time.Sleep(time.Millisecond * 400) // Имитация работы
	}
	close(out)
}

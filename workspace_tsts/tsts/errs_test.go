package tsts

import (
	"errors"
	"testing"
)

// 1. Создаем понятный эталон ошибки (глобально на весь пакет)
var ErrDivisionByZero = errors.New("division by zero")

// 2. Сама функция. Теперь она возвращает наш четкий эталон
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero // Возвращаем конкретно эталон
	}
	return a / b, nil
}

// 3. Тест, который проверяет функцию «на вшивость»
func TestDivide(t *testing.T) {
	_, err := Divide(1, 0) // Специально ломаем логику

	// Проверка 1: А была ли вообще ошибка?
	if err == nil {
		t.Fatal("Divide(1, 0) expected error, got nil") // Fatal сразу остановит тест, если ошибки нет
	}

	// Проверка 2: Проверка на какую-либо другую ошибку (например, падение БД)
	if !errors.Is(err, ErrDivisionByZero) {
		t.Errorf("Ожидали ошибку '%v', но вместо неё получили: '%v'", ErrDivisionByZero, err)
	}
}

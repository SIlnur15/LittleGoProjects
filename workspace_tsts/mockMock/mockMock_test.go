package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 3. Структура-мок, реализующая интерфейс DB
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Get(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

// 4. Сам unit-тест
func TestService(t *testing.T) {
	mockDB := new(MockDB)

	// Настраиваем ожидание: при вызове Get("test") вернем ("value", nil)
	mockDB.On("Get", "test").Return("value", nil)

	// Вызываем тестируемую функцию, передавая мок вместо реальной БД
	result, err := MyService(mockDB, "test")

	// Проверяем результат
	assert.NoError(t, err)
	assert.Equal(t, "value", result)

	// Убеждаемся, что все ожидаемые вызовы (On) действительно произошли
	mockDB.AssertExpectations(t)
}

package tsts_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUserAPI проверяет интеграцию API + сервис + БД
func TestUserAPI(t *testing.T) {
	t.Parallel()

	// 1. Мокаем базу данных
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// 1.5. Настраиваем ожидания для SQL-запроса
	columns := []string{"id", "name"}
	mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "John Doe"))

	// 2. Инициализируем ваш сервис (подставьте вашу реальную функцию)
	svc := NewUserService(db)

	// 3. Запускаем тестовый HTTP-сервер (подставьте ваш реальный обработчик)
	ts := httptest.NewServer(NewHandler(svc))
	defer ts.Close()

	// 4. Выполняем HTTP-запрос к тестовому серверу
	resp, err := http.Get(ts.URL + "/users/1")
	require.NoError(t, err)
	defer resp.Body.Close()

	// 5. Проверяем статус ответа
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 6. Проверяем, что все запланированные SQL-запросы были вызваны
	require.NoError(t, mock.ExpectationsWereMet())
}

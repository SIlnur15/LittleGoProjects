package main

// 1. Интерфейс, который мы хотим замокировать
type DB interface {
	Get(key string) (string, error)
}

// 2. Реализуем функцию, которую будем тестировать
func MyService(db DB, key string) (string, error) {
	val, err := db.Get(key)
	if err != nil {
		return "", err
	}
	return val, nil
}

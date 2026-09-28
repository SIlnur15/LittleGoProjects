package main

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// scanRow сканирует текущую строку rows в структуру dest с помощью рефлексии
func scanRow(rows *sql.Rows, dest interface{}) error {
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	val := reflect.ValueOf(dest).Elem()         // Получаем Value структуры
	fields := make([]interface{}, len(columns)) // Создаём срез указателей для Scan

	// Сопоставляем колонки с полями структуры по тегам db или по имени
	for i, col := range columns {
		field := val.FieldByNameFunc(func(fieldName string) bool {
			f, _ := val.Type().FieldByName(fieldName)
			tag := f.Tag.Get("db")
			return tag == col || f.Name == col
		})
		if field.IsValid() && field.CanAddr() {
			fields[i] = field.Addr().Interface()
		} else {
			// Если поле не найдено, создаём заглушку для пропуска значения
			var dummy interface{}
			fields[i] = &dummy
		}
	}

	return rows.Scan(fields...)
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := scanRow(rows, &user); err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", user)
	}
}

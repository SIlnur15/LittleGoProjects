package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Импорт драйвера MySQL
)

func printUsers(db *sql.DB) error {
	rows, err := db.Query(
		"SELECT id, name FROM users WHERE active = ?",
		true,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			return err
		}

		fmt.Println(id, name)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := printUsers(db); err != nil {
		log.Fatal(err)
	}
}

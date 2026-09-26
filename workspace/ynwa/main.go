package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func listDrivers() {
	for _, driver := range sql.Drivers() {
		fmt.Printf("Driver: %v\n", driver)
	}
}

func main() {
	// db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	listDrivers()
}

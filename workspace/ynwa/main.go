package main

import (
	"encoding/json"
	"os"
)

// Пример структуры
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Чтение структуры из JSON файла
func loadJSON(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

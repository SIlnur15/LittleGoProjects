package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Msg  string
	Code int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s (код: %d)", e.Msg, e.Code)
}

func main() {
	err := doSomething()

	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Printf("Ошибка: %v, код: %d\n", myErr.Msg, myErr.Code)
	}
}

func doSomething() error {
	return &MyError{Msg: "ошибка сервера", Code: 500}
}

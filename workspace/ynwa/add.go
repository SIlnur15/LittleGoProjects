package main

import "fmt"

type Reader interface {
	Read(p []byte) (n int, err error)
}

type LoggingReader struct {
	r Reader
}

func (lr *LoggingReader) Read(p []byte) (int, error) {
	fmt.Println("Начинаем читать данные")
	n, err := lr.r.Read(p)
	fmt.Println("Прочитано байт:", n)
	return n, err
}

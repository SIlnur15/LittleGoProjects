package main

import (
	"bytes"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("grep", "go")

	var stdin bytes.Buffer // Создаем и записываем данные в STDIN команды
	cmd.Stdin = &stdin
	stdin.Write([]byte("golang\npython\njava\n"))

	var stdout, stderr bytes.Buffer // Буферы для STDOUT и STDERR
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run() // Запускаем команду
	if err != nil {
		log.Fatalf("Ошибка выполнения: %s\nStderr: %s", err, stderr.String())
	}

	log.Printf("Stdout: %s", stdout.String())
}

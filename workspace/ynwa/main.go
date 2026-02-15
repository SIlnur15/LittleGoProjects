package main

import (
	"fmt"
	"os"
)

func main() {
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, entry := range entries {
		info, _ := entry.Info() // Получаем os.FileInfo (если нужно)
		fmt.Printf(
			"Name: %-10s | IsDir: %-5t | Size: %d\n",
			entry.Name(),
			entry.IsDir(),
			info.Size(),
		)
	}
}

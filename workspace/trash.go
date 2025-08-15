package main

import (
	"fmt"
	"time"
)

func main() {
	// считываем временной отрезок из os.Stdin
	// гарантируется, что значение корректное
	var s string
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)

	// выведите исходное значение (s) и количество минут в нем в формате "исходное = X min"
	// используйте метод .Minutes() объекта d
	fmt.Printf("%s = %v min", s, d.Minutes())
}

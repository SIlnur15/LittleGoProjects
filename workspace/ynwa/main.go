package main

import (
	"fmt"
	"log"
	"sort"
	"sync"
)

// Структура для хранения информации об устройстве
type Device struct {
	Name string
	IP   string
	Data string
}

// Форматированный вывод
func (d Device) String() string {
	return fmt.Sprintf("Device{Name: %q, IP: %q, Data: %q}", d.Name, d.IP, d.Data)
}

// Эмуляция SSH-команды: возвращает данные как строку
func sshRunCommand(dev Device, user, password, cmd string) (string, error) {
	return fmt.Sprintf("Данные по устройству %s", dev.Name), nil
}

func main() {
	// Создаём срез указателей на Device
	devices := []*Device{
		{"Device1", "192.168.1.1", ""},
		{"Device2", "192.168.1.2", ""},
		{"Device3", "192.168.1.3", ""},
	}
	user := "admin"
	password := "password"
	command := "show version"
	maxConcurrent := 3

	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrent)

	for _, dev := range devices {
		wg.Add(1)
		go worker(dev, user, password, command, sem, &wg)
	}

	wg.Wait()

	// Сортируем устройства по IP
	sort.Slice(devices, func(i, j int) bool {
		return devices[i].IP < devices[j].IP
	})

	// Выводим результаты
	fmt.Println("\n--- Итоговые результаты (отсортировано по IP) ---")
	for _, dev := range devices {
		fmt.Println(dev)
	}
}

// Рабочая функция: принимает указатель на устройство
func worker(
	device *Device, // ← указатель! Первая пропущенная строка
	user, password, cmd string,
	sem chan struct{},
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	sem <- struct{}{} // захват семафора
	defer func() { <-sem }()

	output, err := sshRunCommand(*device, user, password, cmd) // Вторая пропущенная строка
	if err != nil {
		log.Printf("Ошибка на устройстве %s: %v", device.IP, err)
		return
	}

	// Обновляем поле Data — безопасно, так как только одна горутина работает с этим устройством
	device.Data = output // Третья пропущенная строка
}

package main

import (
	"fmt"
)

func main() {
	var numberToScan int
	fmt.Scan(&numberToScan)

	// Извлекаем цифры
	hundreds := numberToScan / 100   // Сотни
	tens := (numberToScan / 10) % 10 // Десятки
	units := numberToScan % 10       // Единицы

	// Проверяем, все ли цифры различны
	if hundreds == tens || hundreds == units || tens == units {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

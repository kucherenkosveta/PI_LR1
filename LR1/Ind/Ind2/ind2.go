package main

import "fmt"

func main() {
	var monitorCost float64 = 8000
	var systemBlockCost float64 = 60000
	var keyboardCost float64 = 1500
	var mouseCost float64 = 800
	var numberOfComputers int

	fmt.Print("Введите количество компьютеров: ")
	fmt.Scanln(&numberOfComputers)

	// Вычисление общей стоимости компьютеров
	totalCost := float64(numberOfComputers) * (monitorCost + systemBlockCost + keyboardCost + mouseCost)

	fmt.Println("Общая стоимость", numberOfComputers, "компьютеров: ", totalCost)
}

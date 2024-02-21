package main

import (
	"fmt"
	"math"
)

func main() {
	// Задаем радиус шара
	var radius float64 = 5

	// Вычисляем объем шара
	V := (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)

	// Вычисляем площадь поверхности шара
	S := 4 * math.Pi * math.Pow(radius, 2)

	fmt.Println("Объем шара: ", V)
	fmt.Println("Площадь поверхности шара: ", S)
}

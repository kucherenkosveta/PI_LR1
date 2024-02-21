package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 8  // полуось эллипса по оси х
	var b float64 = 10 // полуось эллипса по оси у

	// Функция для вычисления y в зависимости от x
	y := func(x float64) float64 {
		return b * math.Sqrt(1-math.Pow(x/a, 2))
	}

	// Шаг интегрирования
	step := 0.01

	// Вычисление площади поверхности
	S := 0.0
	for x := -a; x < a; x += step {
		rectWidth := step
		rectHeight := y(x)
		S += rectWidth * rectHeight
	}

	// Вычисление объема
	V := 0.0
	for x := -a; x < a; x += step {
		rectWidth := step
		rectHeight := math.Pow(y(x), 2)
		V += rectWidth * rectHeight
	}

	fmt.Println("Площадь поверхности: ", S)
	fmt.Println("Объем тела: ", V)
}
